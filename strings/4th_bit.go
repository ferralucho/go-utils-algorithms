package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the 'fourthBit' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER number as parameter.
 */

func fourthBit(number int32) int32 {
	// Write your code here
	binary := strconv.FormatInt(int64(number), 2)
	fmt.Printf("Binary string is %s\n", binary)

	toExtract := len(binary) - 4

	fmt.Printf("Binary is %s\n", string(binary[toExtract]))

	r := string(binary[toExtract])

	i, _ := strconv.ParseInt(r, 10, 32)
	fmt.Printf("Int is %d\n", i)
	return int32(i)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	numberTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	number := int32(numberTemp)

	result := fourthBit(number)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
}

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
