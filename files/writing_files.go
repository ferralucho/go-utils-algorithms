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
	"os"
	"strconv"
	"strings"
)

func writeToFile(bytesChannel chan []byte, doneChannel chan bool, errChannel chan error) {
	file, err := os.Create("writing_file.txt")
	errChannel <- err
	if err != nil {
		return
	}
	for {
		select {
		case b := <-bytesChannel:

			fmt.Printf("%T, %v\n", b, b)
			_, err := file.Write(b)
			fmt.Printf("Sender: Write %s value", b)

			if err != nil {
				errChannel <- err
			}
		case l := <-doneChannel:
			fmt.Printf("Sender: Done %s value\n", l)

		}
	}
}

func main() {
	//reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	//sizeByteArr := readerLine(reader)
	intSize, err := strconv.Atoi("2")

	if err == nil {
		fmt.Printf("%T, %v\n", intSize, intSize)
	}

	bytesChannel := make(chan []byte, intSize)
	errChannel := make(chan error)
	doneChannel := make(chan bool, 1)

	defer close(bytesChannel)
	defer close(errChannel)
	defer close(doneChannel)

	go func() {

		for {
			select {
			case l := <-errChannel:
				fmt.Printf("Sender: Error %s value\n", l)

			}
		}

	}()

	for i := 0; i < intSize; i++ {
		in := "primero" + strconv.Itoa(i)
		bytesChannel <- []byte(in + "\n")
	}

	writeToFile(bytesChannel, doneChannel, errChannel)
	doneChannel <- true
}

func readerLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()

	if err == io.EOF {
		return ""
	}
	return strings.TrimRight(string(str), "\r\n")
}
