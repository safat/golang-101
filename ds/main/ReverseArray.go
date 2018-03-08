package main

import "fmt"
import "strconv"

func main() {
	var output string
	var n int

	fmt.Scan(&n)
	a := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])

		if i > 0 {
			output = " " + output
		}

		output = strconv.Itoa(a[i]) + output
	}

	fmt.Println(output)
}

/*
	Problem Description
	Print all  integers in  in reverse order as a single line of space-separated integers.

	https://www.hackerrank.com/challenges/arrays-ds/problem
*/

/*
	Solution 2

	Actually reverse the array like below and print out the result

	   for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}

	Reference: https://stackoverflow.com/questions/19239449/how-do-i-reverse-an-array-in-go

 */
