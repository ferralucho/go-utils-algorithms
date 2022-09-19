/*
There is an array of strings. Write a function that will:
1. create a file
2. send nil to errChannel as a signal to the main function that it can begin sending data
3. accept strings from an array converted to byte arrays passed from the main function
4. write the strings to the file

This should be done using a minimal amount of memory allocations.
The name of the file to write is stored in a global variable filename.

Example

inputArray = ["Lorem ", "ipsum ", "dolor ", "sit ", "amet"]
The resulting file should contain "Lorem ipsum dolor sit amet".

Function Description

Complete the function writeToFile in the editor below. The function must be void.

writeToFile has the following parameters:
bytesChannel chan []byte: a channel for receiving bytes from the main function for writing them to the
file.
doneChannel chan bool: a channel for receiving the signal from the main function that all data is sent.
errChannel chan error: a channel through errors will be sent to the main function (including nil).

Constraints
The length of the input array does not exceed 1000.

Input Format For Custom Testing

The first line contains an integer, m, denoting the size of the byte array which should be read from the
file at every iteration.
The second line contains an integer, n, denoting m + the number of bytes which should be skipped after
reading at every iteration.
The third line contains a string inputString denoting the contents of the file.

Sample Case 0----------

Sample Input For Custom Testing
2
Hello
World

Explanation
2 strings are sent to the writeToFile function which writes them to the file, then the main function will
read them from this file where they get joined.


Sample case 1----------

Sample Input For Custom Testing
5
Lorem
ipsum
dolor
sit
amet

Sample Output
Lorem ipsum dolor sit amet

Explanation
5 strings are sent to the writeToFile function which writes them to the file. The main function then reads
them from this file where they get joined.

*/
package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
)

/*
 * Complete the 'writeToFile' function below.
 *
 * The function accepts following parameters:
 *  1. chan []byte bytesChannel
 *  2. chan bool doneChannel
 *  3. chan error errChannel
 */

func writeToFile(bytesChannel chan []byte, doneChannel chan bool, errChannel chan error) {
	file, err := os.Create(filename)
	errChannel <- nil
	if err != nil {
		return
	}

	for {
		select {
		case val, ok := <-bytesChannel:
			if !ok {
				fmt.Printf("Bytes channel closed")
				return
			}
			fmt.Printf("%v\n", val)
			_, err := file.Write(val)

			if err != nil {
				errChannel <- err
			}
			errChannel <- nil
		case d := <-doneChannel:
			fmt.Printf("Done %t value\n", d)
			return
		case e := <-errChannel:
			fmt.Printf("Error %t value\n", e)
		}
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	inputArrayCount, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	var inputArray []string

	for i := 0; i < int(inputArrayCount); i++ {
		inputArrayItem := readLine(reader)
		inputArray = append(inputArray, inputArrayItem)
	}

	bytesChannel, doneChannel, errChannel := make(chan []byte), make(chan bool), make(chan error)
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	allocBefore := ms.Alloc
	go writeToFile(bytesChannel, doneChannel, errChannel)
	err = <-errChannel
	if err != nil {
		panic(err)
	}
	for _, b := range inputArray {
		bytesChannel <- []byte(b)
		err := <-errChannel
		if err != nil {
			fmt.Fprintf(writer, "Critical error: %s", err.Error())
			break
		}
	}
	doneChannel <- true
	runtime.ReadMemStats(&ms)
	allocAfter := ms.Alloc
	fmt.Printf("Total memory allocated: %d bytes\n", allocAfter-allocBefore)
	if allocAfter-allocBefore > 10000 {
		fmt.Fprintf(writer, "Too much memory allocated, maximum 10000 bytes needed")
	} else {
		b, err := ioutil.ReadFile(filename)
		if err == nil {
			fmt.Fprintf(writer, "%s\n", string(b))
		} else {
			fmt.Fprintf(writer, "Critical error: %s", err.Error())
		}
	}

	writer.Flush()
}

const filename = "output"

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
