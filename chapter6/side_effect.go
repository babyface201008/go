package main

import "fmt"

func main() {
	n := 0
	relpy := &n
	Multiply(10, 5, relpy)
	fmt.Println("Multiply:", *relpy)
}

func Multiply(a, b int, reply *int) {
	*reply = a * b
}
