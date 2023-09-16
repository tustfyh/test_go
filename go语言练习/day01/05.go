package main

import "fmt"

//建立一个 Go 程序打印下面的内容（到 100 个字符）：
/*  A
AA
AAA
AAAA
...
*/
func test01() {
	for i := 1; i < 101; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("A")
		}
		print("\n")
	}
}
func test02() {
	a := "A"
	for i := 0; i < 100; i++ {
		fmt.Printf("%v\n", a)
		a = a + "A"
	}

}
func main() {
	test01()
}
