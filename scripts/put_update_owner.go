package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

type Record struct {
	LoanId                    string
	RestructureSuccess        bool
	RestructureDescription    string
	RestructureStatusResponse int
}

var cName string //write-java-beta_credits-api.furyapps.io
var furyToken string

func restructureCreditLine(record *Record) error {
	record.RestructureDescription = fmt.Sprintf("Executing restructure ownership loan in Core %s", record.LoanId)

	client := &http.Client{}
	url := fmt.Sprintf("https://%s/credits/loans/owner/%s", cName, record.LoanId)

	bodyMap := make(map[string]interface{})

	body, _ := json.Marshal(bodyMap)
	req, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Caller-Scopes", "admin")
	req.Header.Set("x-auth-token", furyToken)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Update error - ", url)
		record.RestructureDescription = fmt.Sprintf("Update error - %s", url)

		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		record.RestructureSuccess = true
		record.RestructureDescription = fmt.Sprintf("Update restructure loan ownership Success - %d - %s", resp.StatusCode, url)

	} else {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		record.RestructureDescription = fmt.Sprintf("Update restructure loan ownership Error - %d - %s - %s", resp.StatusCode, url, string(bodyBytes))
	}
	record.RestructureStatusResponse = resp.StatusCode

	return nil
}

func worker(id int, records <-chan *Record, results chan<- *Record, group *sync.WaitGroup) {
	fmt.Println("worker", id, "started")
	defer func() {
		fmt.Println("worker", id, "Deferred func")
		group.Done()
	}()
	for record := range records {
		fmt.Println("worker", id, "processing record ", record.LoanId)
		err := restructureCreditLine(record)
		if err != nil {
			fmt.Println("worker", id, "error on record:", record.LoanId)
			fmt.Println(err)
		}
		fmt.Println("worker", id, "finished record", record.LoanId)
		results <- record
	}
	fmt.Println("worker", id, "done")
}

func parallelRead(reader *csv.Reader, records chan *Record, results chan *Record) {
	fmt.Println("Start reading file")
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		record := Record{
			LoanId: line[0],
		}

		records <- &record
	}
	fmt.Println("Finished reading file")
	close(records)
}

func main() {
	inputPtr := flag.String("input", "input.csv", "The input file to process")
	hasHeaderPtr := flag.Bool("hasHeader", true, "Param indicating if the input file has a header or not")
	routinesPtr := flag.Int("routines", 1, "The amount of routines to use for processing the file")
	url_flag := flag.String("url", "", "Indicate the url to perform the put.")
	token_flag := flag.String("token", "", "Fury token to use.")

	startTime := time.Now()

	flag.Parse()

	cName = *url_flag
	furyToken = *token_flag

	fmt.Println(*inputPtr)
	fmt.Println(*hasHeaderPtr)
	fmt.Println(*routinesPtr)
	fmt.Println(*url_flag)
	fmt.Println(*token_flag)

	//------------------------------Lectura y carga en array de datos de Record --------------------------------------//
	// START FILE HANDLING
	inputFile, err := os.Open(*inputPtr)
	if err != nil {
		log.Fatal("Error creating input file", err)
		return
	}

	outputFile, err := os.Create("./result.csv")
	if err != nil {
		log.Fatal("Error creating result file")
		return
	}
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	// Create a new reader.
	reader := csv.NewReader(bufio.NewReader(inputFile))
	if *hasHeaderPtr {
		_, err := reader.Read()
		if err != nil {
			log.Fatal("Error reading header")
			return
		}
		header := []string{"LoanId", "RestructureSuccess", "RestructureStatusResponse", "RestructureDescription"}
		err = writer.Write(header)
		if err != nil {
			log.Fatal("Error writting header")
			return
		}
	}
	// END FILE HANDLING

	var successCounter int64
	var failCounter int64

	// CREACION CHANNELS Y WORKERS
	records := make(chan *Record, 100)
	results := make(chan *Record, 100)

	group := sync.WaitGroup{}
	group.Add(*routinesPtr)
	for w := 1; w <= *routinesPtr; w++ {
		go worker(w, records, results, &group)
	}
	// END CREACION CHANNELS Y WORKERS
	// Cerrar results cuando no se hayan consumido más cosas en los workers
	go func() {
		group.Wait()
		close(results)
	}()

	go parallelRead(reader, records, results)
	count := 0
	fmt.Println("Starting to wait for results")
	for record := range results {
		count++

		err = writer.Write([]string{record.LoanId, strconv.FormatBool(record.RestructureSuccess), strconv.Itoa(record.RestructureStatusResponse), record.RestructureDescription})
		if err != nil {
			fmt.Println(fmt.Sprintf("Error writting record: %s", record.LoanId))
		}
		fmt.Println(fmt.Sprintf("Wrote record: %s", record.LoanId))
		if count%100 == 0 {
			writer.Flush()
		}

		if record.RestructureSuccess {
			successCounter++
		} else {
			failCounter++
		}
		fmt.Println("Processed:", count)
	}

	endTime := time.Now()

	fmt.Println(fmt.Sprintf("Successful records: %d", successCounter))
	fmt.Println(fmt.Sprintf("Failed records: %d", failCounter))
	fmt.Println(fmt.Sprintf("Time elapsed: %f min.", endTime.Sub(startTime).Minutes()))
}
