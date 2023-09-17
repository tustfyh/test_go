// 最小值和最大值
// 编写一个函数，找到 int slice ( []int ) 中的最大值。
// 编写一个函数，找到 int slice ( []int ) 中的最小值。
package main
import "fmt"
func find_min(num []int) int {

	var tmp = num[0]
	for _, v := range num {
		if tmp > v {
			tmp = v
		}
	}
	return tmp
}
func find_max(num []int) int {

	var tmp = num[0]
	for _, v := range num {
		if tmp < v {
			tmp = v
		}
	}
	return tmp
}
func main() {
	arr := []int{1, 5, 6, 85, 12, 36, 95, -5}
	fmt.Printf("min=%v\n", find_min(arr))
	fmt.Printf("max=%v\n", find_max(arr))
}
