package main

import "fmt"

// 创建一个固定大小保存整数的栈。它无须超出限制的增长。
// 定义 push 函数—— 将数据放入栈，和 pop 函数——从栈中取得内容。栈应当是后进先出（LIFO） 的。
type Stack struct {
	item []int
}

func Push(e int, S *Stack) {
	S.item = append(S.item, e)
}
func Pop(S *Stack) (ret int) {
	length := len(S.item)
	ret = S.item[length-1]
	S.item = S.item[:length-1]
	return ret
}
func main() {
	s := &Stack{}
	for i := 0; i < 10; i++ {
		Push(i, s)
	}
	for i := 0; i < 10; i++ {

		fmt.Printf("%v\n", Pop(s))
	}
}
