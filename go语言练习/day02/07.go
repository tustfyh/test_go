// 函数返回一个函数
// 编写一个函数返回另一个函数，返回的函数的作用是对一个整数 +2。函数的名 称叫做 plusTwo 。
// 然后可以像下面这样使用： p := plusTwo() fmt.Printf("%v\n", p(2)) 应该打印 4。
package main

import "fmt"
//回调函数
func test(a int) int {
	return a + 2
}
func plustwo() func(int2 int) int {

	return test
}
func main() {
	p := plustwo()
	fmt.Printf("%v ", p(2))
}
