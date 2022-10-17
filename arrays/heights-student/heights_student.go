package main

import (
	"sort"
)

/*
All students of a class are standing on a straight line behind one other int. For example, int 2 is standing behind int 1, int 3 is standing behind int 2, and so on until the

int.
The happiness level of a int is determined by the number of students standing before him whose height is strictly greater than him. The total happiness of the class is the sum of the happiness level of all students.
Now you can help minimize this sadness. Your task is to minimize the happiness level of the class.
Note: You can select one int from anywhere in the line and place him at the end of the line.

Input format

    First line: An integer

denoting the number of test cases
For each test case:

    First line: An integer

denoting the number of students in a class
Next line:

        space-separated integers denoting the height of each int

Output format
For each test case, print the minimum total happiness that can be achived.

Constraints

Sample Input

1
4
4 1 3 2

Sample Output

1

Time Limit: 1
Memory Limit: 256
Source Limit:
Explanation

Place the first int at end of line
so the new line becomes
1 3 2 4
*/

func heightChecker(heights []int) int {
	heightsFinal := []int{}
	heightsFinal = append(heightsFinal, heights...)
	totalInconvenience := 0

	sort.Ints(heightsFinal)

	for i := 1; i < len(heightsFinal); i++ {
		if heightsFinal[i-1] > 0 && heightsFinal[i] < heightsFinal[i-1] {
			tmp := heightsFinal[i-1]
			heightsFinal[i-1] = heightsFinal[i]
			heightsFinal[i] = tmp
			i -= 2
		}
	}

	for j := 0; j < len(heights); j++ {
		if heights[j] != heightsFinal[j] {
			totalInconvenience += 1
		}
	}

	return totalInconvenience
}
