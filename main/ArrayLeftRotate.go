package main

import "fmt"
import "bytes"
import "strconv"

func main() {
	var n, d int
	var buffer bytes.Buffer

	fmt.Scan(&n, &d)

	var numbers = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&numbers[i])
	}

	for i := d; i < n+d; i++ {
		buffer.WriteString(strconv.Itoa(numbers[i%n]))

		if i < n+d {
			buffer.WriteString(" ")
		}
	}

	fmt.Println(buffer.String())
}

//https://www.hackerrank.com/challenges/array-left-rotation/problem