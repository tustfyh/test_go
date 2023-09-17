// 编写一个针对 int 类型的 slice 冒泡排序的函数
package main

import "fmt"

func Bubble_Sort(num []int) []int {

	lenght := len(num)
	var tmp int
	for i := 0; i < lenght-1; i++ {
		for j := 0; j < lenght-1-i; j++ {
			if num[j] > num[j+1] {
				tmp = num[j]
				num[j] = num[j+1]
				num[j+1] = tmp
			}
		}
	}
	return num
}

func main() {
	arr := []int{85, 41, 25, 96, 455, 488, 3, 25}
	Bubble_Sort(arr)
	for _, v := range arr {
		fmt.Printf("%v ", v)
	}
}
