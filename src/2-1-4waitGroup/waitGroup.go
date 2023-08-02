package main

import (
	"sync"
)

func main() {
	// 新建一个WaitGroup结构
	wg := sync.WaitGroup{}

	// 计数器 + 5
	wg.Add(5)

	// 并发快速打印 0-5
	for i := 0; i < 5; i++ {
		go func(num int) {
			// 延迟计数器 -1
			defer wg.Done()
			println(num)
		}(i)
	}

	// 阻塞等待计数器清零
	wg.Wait()
}
