package main

import "fmt"

func main() {

	//slice1 := make([]int, 0, 10)

	/*for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("The length of slice is %d\n", len(slice1))
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	*/
	var ar = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var a = ar[5:7]
	for i := 0; i < len(a); i++ {
		fmt.Printf("Slice is %d\n", a[i])
	}

	a = a[0:4]
	for i := 0; i < len(a); i++ {
		fmt.Printf("Slice is %d\n", a[i])
	}
}
