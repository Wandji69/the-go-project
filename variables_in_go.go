package main

import "fmt"

func main() {
	var x int
	var A, B float64
	var c = 7
	var d string
	var e, f, g = true, 2.3, "four"
	x = 10
	A = 2.5
	B = 3.5
	d = "hello"
	fmt.Printf("x = %d, A = %f,B = %f, d = %s, e = %v, f = %f, g = %s, c = %d \n", x, A, B, d, e, f, g, c)
}

// the above prints: x = 10, A = 2.500000,B = 3.500000, d = hello, e = true, f = 2.300000, g = four, c = 7
