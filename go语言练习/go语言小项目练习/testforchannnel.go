// 关于协程的练以及WaitGroup的使用
package main

import (
	"fmt"
	"sync"
)
func accumulation(n int) uint64 {
	var tmp uint64
	for i := 1; i <= n; i++ {
		tmp += uint64(i)
	}
	return tmp
}

func writeintchan(intchan chan int) {
	for i := 1; i < 2001; i++ {
		intchan <- i
	}
	close(intchan)
}

func readintchan(intchan chan int, reschan chan uint64, wg *sync.WaitGroup) {
	for v := range intchan {
		reschan <- accumulation(v)
		wg.Done()
	}
}
func main() {

	var wg sync.WaitGroup //值传递
	intchan := make(chan int, 2000)
	reschan := make(chan uint64, 2000)
	wg.Add(2000)
	go writeintchan(intchan)
	for i := 0; i < 8; i++ {
		go readintchan(intchan, reschan, &wg)
	}
	wg.Wait()
	close(reschan)
	for v := range reschan {
		fmt.Println(v)
	}
}
