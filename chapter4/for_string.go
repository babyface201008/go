package main

import "fmt"

func main() {
	s := "\u4f60\u597d\u554a"
	for i, c := range s {
		fmt.Printf("%d:%c\n", i, c)
	}
	fmt.Println(s)
}
