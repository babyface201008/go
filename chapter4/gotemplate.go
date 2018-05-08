package main

import (
	"fmt"
)

//全局变量在本包内可调用

const c = "C"

var v int = 5

func main() { //主函数代码调用地方
	test()
}

func test() {
	fmt.Println(v)
	fmt.Println(c)
}
