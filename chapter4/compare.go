package main

import "fmt"

func main() {
	s1 := []byte("abcdef")
	s2 := []byte("abefre")

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(Compare(s1, s2))
}

func Compare(a, b []byte) int {

	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}

	//数组长度可能不同
	switch {
	case len(a) < len(b):
		return -1
	case len(a) > len(b):
		return 1
	}
	return 0
}
