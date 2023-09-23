package main

import (
	"fmt"
	"sort"
)

func main() {

	var name = [...]int{0, 15, 25, 31, 45, 57, 61, 79}
	sl := make([]int, len(name))
	copy(sl, name[:])
	sl = append(sl, 9)
	sort.Ints(sl)
	arr := []int(sl)
	fmt.Println(arr)
	fmt.Printf("%T", arr)

}
