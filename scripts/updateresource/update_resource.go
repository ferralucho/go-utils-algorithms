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

/*
Resource massive update. Reads from an input CSV,
reads lines and perform requests using workers (could use it to update resources).
Creates an output CSV.
You can set the amount of go routines.

The input.csv must be like this:

resource_id
14806111
14807112
*/

type Record struct {
	ResourceID           string
	UpdateSuccess        bool
	UpdateDescription    string
	UpdateStatusResponse int
}

var cName string
var token string

func updateResource(record *Record) error {
	record.UpdateDescription = fmt.Sprintf("Executing update %s", record.ResourceID)

	client := &http.Client{}
	url := fmt.Sprintf("https://%s/store/resource/%s", cName, record.ResourceID)

	bodyMap := make(map[string]interface{})

	body, _ := json.Marshal(bodyMap)
	req, _ := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Caller-Scopes", "admin")
	req.Header.Set("x-auth-token", token)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Update error - ", url)
		record.UpdateDescription = fmt.Sprintf("Update error - %s", url)

		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		record.UpdateSuccess = true
		record.UpdateDescription = fmt.Sprintf("Update resource Success - %d - %s", resp.StatusCode, url)

	} else {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		record.UpdateDescription = fmt.Sprintf("Update resource Error - %d - %s - %s", resp.StatusCode, url, string(bodyBytes))
	}
	record.UpdateStatusResponse = resp.StatusCode

	return nil
}

func worker(id int, records <-chan *Record, results chan<- *Record, group *sync.WaitGroup) {
	fmt.Println("worker", id, "started")
	defer func() {
		fmt.Println("worker", id, "Deferred func")
		group.Done()
	}()
	for record := range records {
		fmt.Println("worker", id, "processing record ", record.ResourceID)
		err := updateResource(record)
		if err != nil {
			fmt.Println("worker", id, "error on record:", record.ResourceID)
			fmt.Println(err)
		}
		fmt.Println("worker", id, "finished record", record.ResourceID)
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
			ResourceID: line[0],
		}

		records <- &record
	}
	fmt.Println("Finished reading file")
	close(records)
}

func main() {
	inputPtr := flag.String("input", "input.csv", "The input file to process")
	hasHeaderPtr := flag.Bool("hasHeader", true, "Param indicating if the input file has a header or not")
	routinesPtr := flag.Int("routines", 10, "The amount of routines to use for processing the file")
	url_flag := flag.String("url", "localhost:8080", "Indicate the url to perform the put.")
	token_flag := flag.String("token", "", "Token to use.")

	startTime := time.Now()

	flag.Parse()

	cName = *url_flag
	token = *token_flag

	fmt.Println(*inputPtr)
	fmt.Println(*hasHeaderPtr)
	fmt.Println(*routinesPtr)
	fmt.Println(*url_flag)
	fmt.Println(*token_flag)

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

	reader := csv.NewReader(bufio.NewReader(inputFile))
	if *hasHeaderPtr {
		_, err := reader.Read()
		if err != nil {
			log.Fatal("Error reading header")
			return
		}
		header := []string{"ResourceID", "UpdateSuccess", "UpdateStatusResponse", "UpdateDescription"}
		err = writer.Write(header)
		if err != nil {
			log.Fatal("Error writting header")
			return
		}
	}

	var successCounter int64
	var failCounter int64

	records := make(chan *Record, 100)
	results := make(chan *Record, 100)

	group := sync.WaitGroup{}
	group.Add(*routinesPtr)
	for w := 1; w <= *routinesPtr; w++ {
		go worker(w, records, results, &group)
	}
	go func() {
		group.Wait()
		close(results)
	}()

	go parallelRead(reader, records, results)
	count := 0
	fmt.Println("Starting to wait for results")
	for record := range results {
		count++

		err = writer.Write([]string{record.ResourceID, strconv.FormatBool(record.UpdateSuccess), strconv.Itoa(record.UpdateStatusResponse), record.UpdateDescription})
		if err != nil {
			fmt.Println(fmt.Sprintf("Error writting record: %s", record.ResourceID))
		}
		fmt.Println(fmt.Sprintf("Wrote record: %s", record.ResourceID))
		if count%100 == 0 {
			writer.Flush()
		}

		if record.UpdateSuccess {
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
