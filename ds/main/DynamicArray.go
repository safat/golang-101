package main

import (
	"fmt"
	"bytes"
	"strconv"
)

func main() {
	var n, q int
	var buffer bytes.Buffer

	fmt.Scan(&n, &q)

	var lastAnswer = 0
	var listOfSequence = make([][] int, n)

	for i := 0; i < q; i++ {
		var op, x, y int

		fmt.Scan(&op, &x, &y)
		index := (x ^ lastAnswer) % n

		switch op {
		case 1:
			listOfSequence[index] = append(listOfSequence[index], y)

		case 2:
			lastAnswer = listOfSequence[index][y%len(listOfSequence[index])]
			buffer.WriteString(strconv.Itoa(lastAnswer))
			buffer.WriteString("\n")
		}
	}

	fmt.Print(buffer.String())
}

//https://www.hackerrank.com/challenges/dynamic-array/problem