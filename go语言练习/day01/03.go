package main

import "fmt"

// 写这个循环，使其遍历一个 array，并将这个 array 打印到屏幕上。
func test(arr [6]int) {
	for _, v := range arr {
		fmt.Printf("v=%v\n", v)
	}
}
func main() {
	var arr = [6]int{1, 2, 3, 44, 58, 645}
	test(arr)
}
