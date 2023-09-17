package main

import "fmt"

// 编写函数，返回其（两个）参数正确的（自然）数字顺序：从大到小
func test(a int, b int) (int, int) {

	if a > b {
		return a, b
	} else {
		return b, a
	}

}

func main() {

	a, b := 10, 45
	a, b = test(a, b)
	fmt.Printf("%v,%v", a, b)
}
