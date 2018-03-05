package main

import "fmt"
import "bytes"
import "strconv"

func main() {
	var n, q int
	var buff bytes.Buffer

	var wordMap = make(map[string]int)

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var word string
		fmt.Scan(&word)

		wordMap[word]++
	}

	fmt.Scan(&q)

	for i := 0; i < q; i++ {
		var word string
		fmt.Scan(&word)

		buff.WriteString(strconv.Itoa(wordMap[word]))
		buff.WriteString("\n")
	}

	fmt.Print(buff.String())
}

//https://www.hackerrank.com/challenges/sparse-arrays/problem