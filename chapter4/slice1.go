package main

import "fmt"

func main() {

	slice0 := []string{"a", "b", "c", "d", "e"}
	slice1 := slice0[2:4]
	slice2 := slice0[:3]
	fmt.Println(slice0, slice1, slice2)
	slice2[2] = "8"
	fmt.Println(slice0, slice1, slice2)
	fmt.Println(len(slice0))
}
