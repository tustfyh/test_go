// 编写计算一个类型是 float64 的 slice 的平均值的代码。
package main

import "fmt"

func test(arr []float64, length int) float64 {

	var ret float64 = 0
	var tmp float64 = 0
	for _, f := range arr {
		tmp = tmp + f
	}
	l := float64(length)//强制类型转换
	ret = tmp / l

	return ret
}
func main() {

	var arr = []float64{1.01, 1.5, 25.2, 4.1}
	length := len(arr)
	fmt.Printf("平均值为=%v", test(arr, length))
}
