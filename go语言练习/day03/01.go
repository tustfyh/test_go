package main

import (
	"fmt"
	"math/rand"
	"time"
)

func math_name(name [10]float64) (avg float64, max float64, max_num int) {
	tmp01 := 0.0
	max = name[0]
	for i := 0; i < len(name); i++ {
		tmp01 += name[i]
		if max < name[i] {
			max = name[i]
			max_num = i
		}
	}
	avg = tmp01 / float64(len(name))
	return avg, max, max_num
}
func main() {

	var name [10]float64
	fg := 0
	rand.Seed(time.Now().UnixMilli())
	for i := 0; i < 10; i++ {
		name[i] = float64(rand.Intn(100))
		if name[i] == 55 {
			fg = 1
		}

	}
	avg, max, max_num := math_name(name)
	fmt.Printf("avg=%v,max=%v,max_num=%v\n", avg, max, max_num)
	for i := len(name) - 1; i >= 0; i-- {
		fmt.Printf("%v \n", name[i])
	}
	if fg == 1 {
		print("该数组中存在55\n")
	} else {
		print("该数组中不存在55\n")
	}
}
