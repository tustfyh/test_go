package main
//编写函数接受整数类型变参，并且每行打印一个数字。
func test(arr ...int) {
	for _, v := range arr {
		println(v)
	}
}
func main() {
	test(1, 2, 4, 5, 8, 6)
}
