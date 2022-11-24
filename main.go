package main

import (
	"context"
	"fmt"
	"sync"
)

var ch = make(chan int, 1)

func main() {
	var maps = make(map[int]context.CancelFunc)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		ctx, cancelFunc := context.WithCancel(context.Background())
		maps[i] = cancelFunc
		go func(i int, ctx context.Context) {
			select {
			case <-ctx.Done():
				fmt.Println(i, "已接收到")
				wg.Done()
				return
			}
		}(i, ctx)
		go func(int2 int) {
			select {
			case res := <-ch:
				if cancelFunc, ok := maps[res]; ok {
					fmt.Println(res)
					cancelFunc()

				}

			}
		}(i)
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			ch <- i
		}(i)
	}
	wg.Wait()
}
