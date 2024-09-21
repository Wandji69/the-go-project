package main

import "fmt"

func main() {
	// call add function with argument of 10 and 20
	z := add(10, 20)
	// print th evalue of variable z which is the result of add function
	fmt.Println(z)
}

// implementation of function add which takes two arguments and return their sum
func add(x int, y int) int {
	return x + y
}

// this will print 30
