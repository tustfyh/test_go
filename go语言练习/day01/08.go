package main

import "fmt"

//逆序输出字符串

func test(a string, length int) {
	for i := length - 1; i >= 0; i-- {
		fmt.Printf("%c", a[i])
	}
}
func main() {
	a := "footpf"
	length := len(a)
	test(a, length)
}
