package main

import "fmt"

func Sum(x int, y int) int {
	return x + y
}

func Multiply(x int, y int) int {
	return x * y
}

func Sub2(x int, y int) int {
	return x - y
}

func Sub3(x int, y int, z int) int {
	return x - y
}

func main() {
	fmt.Printf("result : %d\n", Sum(5, 6))
}
