package main
import "fmt"

// 解决这个叫做 Fizz-Buzz 的问题： 编写一个程序，打印从 1 到 100 的数字。
// 当是3个倍数数就打印 “Fizz” 代替数字，当是5的倍数就打印 “Buzz” 。当数字同时是3和5的倍数 时，打印 “FizzBuzz” 。
func test() {
	for i := 1; i < 101; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Printf("FizzBuzz\n")
			} else {
				fmt.Printf("Fizz\n")
			}
		} else if i%5 == 0 {
			if i%3 == 0 {
				fmt.Printf("FizzBuzz\n")
			} else {
				fmt.Printf("Buzz\n")
			}
		} else {
			fmt.Printf("%v\n", i)
		}
	}
}
func main() {
	test()
}
