package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

/*
* Practice codes while reading https://blog.golang.org/go-slices-usage-and-internals
*/

func main() {
	x := []int{2, 4, 5}
	y := x

	y[0] = -10

	fmt.Println(x[0])

	// not same
	println("&x: ", &x)
	println("&y: ", &y)

	// same
	println("&x[0]: ", &x[0])
	println("&y[0]: ", &y[0])

	p := [3]string{"Лайка", "Белка", "Стрелка"}
	q := p[:] // a slice referencing the storage of p

	println("&p[0]", &p[0])
	println("&q[0]", &q[0])

	q = append(q, "test")

	// Now that I have appended few more values to q, let's check if the interal array reference
	// got changed to a new one
	println("after a value to slice")
	println("&p[0]", &p[0])
	println("&q[0]", &q[0])

	q = append(q, "test2")

	// Now that I have appended few more values to q, let's check if the interal array reference
	// got changed to a new one
	println("again after appending a value to slice")
	println("&p[0]", &p[0])
	println("&q[0]", &q[0])

	for i := 0; i < 100; i++ {
		q = append(q, string(i))
	}

	println("again after appending 100 values to slice")
	println("&p[0]", &p[0])
	println("&q[0]", &q[0])

	b := CopyDigits("/home/shakhawat.hossain/Downloads/big.txt")

	fmt.Println(len(b))
}

var digitRegexp = regexp.MustCompile("[0-9]+")

func CopyDigits(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = digitRegexp.Find(b)
	c := make([]byte, len(b))

	return append(c, b...)
}
