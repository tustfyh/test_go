package main

//用 goto 改写 问题1 的循环。关键字 for 不可使用。
import (
	"fmt"
)

func test() {
	i := 1
myflag:
	fmt.Printf("i=%v\n", i)
	if i < 10 {
		i++
		goto myflag
	}

}
func main() {
	test()
}
