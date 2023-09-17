package main

import "fmt"

//斐波那契数列以：1,1,2,3,5,8,13,... 开始。或者用数学形式表达：x 1 = 1;x 2 = 1;x n = x n−1 + x n−2 ∀n > 2。
//编写一个接受 int 值的函数，并给出这个值得到的斐波那契数列。

func test(a int) {
	x1 := 1
	x2 := 1
	tmp := 2
	var arr = []int{1, 1}
	if a > 2 {
		for i := 2; i < a; i++ {
			tmp = x1 + x2
			x1 = x2
			x2 = tmp
			arr = append(arr, tmp)
		}
	} else {
		if a == 1 {
			print(1)
		} else {
			fmt.Printf("1,1")
		}
	}
	if a > 2 {
		for i, v := range arr {
			if i != 0 {
				fmt.Printf(",%v",v)
			} else {
				fmt.Printf("%v", v)
			}

		}
	}
}

func main() {
	test(10)
}
