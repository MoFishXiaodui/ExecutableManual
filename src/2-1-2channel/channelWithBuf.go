package main

import "fmt"

func CalSquare() {
	// 1-创建一个无缓冲的channel: src，创建一个有缓冲的channel: dest
	src := make(chan int)
	dest := make(chan int, 3)

	// 2-创建一个立即执行的匿名routine，依次往src管道里面放入数字 0~9
	go func() {
		defer close(src)
		for i := 0; i < 10; i++ {
			src <- i
		}
	}()

	// 3-创建一个立即执行的匿名routine，接收src管道的数字，并计算平方，放入dest管道中
	go func() {
		defer close(dest)
		for i := range src {
			dest <- i * i
		}
	}()

	// 4-接收dest管道的数字并打印出来
	for i := range dest {
		fmt.Println(i)
	}
}

func main() {
	CalSquare()
}
