package main

/*
Implement a function that takes in two input readers (io.Reader) where each contains a string. Merge them as follows.
The two readers should be read one character byte at a time alternately. Push the characters to the resulting string's reader.
If the two strings are of unequal lengths, the longer of the two strings should be trimmed to the length of the shorter string.

Function Description
r1 reads ABCDE from STDIN
r2 reads abcde from STDIN
The resulting string AaBbCcDdEe.

Constraints
1 ≤ length of string ≤ 2000

Sample Input For Custom Testing
abc
ABC
Sample Output
aAbBcC

The readers are merged, alternately taking characters from r1 and r2.

Sample Input For Custom Testing
Hello
World!!!

Sample Output
HWeolrllod
*/

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

/*
* Complete the 'MergeReaders' function below.
*
* The function is expected to return an io.Reader and an error.
* The function accepts following parameters:
* 1. io.Reader r1
* 2. io.Reader r2
 */

func MergeReaders(r1, r2 io.Reader) (io.Reader, error) {
	r1bytes, err1 := ioutil.ReadAll(r1)
	r2bytes, err2 := ioutil.ReadAll(r2)

	if err1 != nil {
		return strings.NewReader(""), err1
	}

	if err2 != nil {
		return strings.NewReader(""), err2
	}
	str1 := string(r1bytes)
	str2 := string(r2bytes)
	len1 := len(str1)
	len2 := len(str2)
	len3 := 0

	if len1 > len2 {
		len3 = len2
	} else if len2 > len1 {
		len3 = len1
	} else {
		len3 = len1
	}

	result := ""

	for i := 0; i < len3; i++ {
		result += str1[i:i+1] + str2[i:i+1]
	}
	return strings.NewReader(result), nil
}
func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)
	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)
	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)
	str1 := readLine(reader)
	str2 := readLine(reader)

	if len(str1) > len(str2) {
		str1 = str1[0:len(str2)]
	} else if len(str2) > len(str1) {
		str2 = str2[0:len(str1)]
	}

	r1 := strings.NewReader(str1)
	r2 := strings.NewReader(str2)
	resultReader, err := MergeReaders(r1, r2)

	if err != nil {
		panic(err)
	}
	result, err := ioutil.ReadAll(resultReader)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(writer, "%s\n", string(result))
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
