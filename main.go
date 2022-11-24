package main

import (
	"fmt"
	"sync"
	"time"
)

var ch = make(chan int, 1)

func main() {

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(int2 int) {
			select {
			case res := <-ch:
				go Test(res)
				wg.Done()
				return
			}
		}(i)
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			ch <- i
		}(i)
		time.Sleep(time.Second)
	}
	wg.Wait()
}
func Test(i int) {
	fmt.Println("已收到", i)
}
