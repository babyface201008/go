package main

import "fmt"

func main() {

	slice1 := make([]int32, 5, 8)
	slice2 := make([]int32, 9)
	slice3 := []int32{}
	slice4 := []int32{1, 2, 3, 4, 5}
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(slice4)
}
