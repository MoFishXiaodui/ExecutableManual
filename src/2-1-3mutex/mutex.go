package main

import (
	"sync"
	"time"
)

var x int64
var lock sync.Mutex

func addWithLock() {
	for i := 0; i < 2000; i++ {
		lock.Lock()
		x += 1
		lock.Unlock()
	}
}

func addWithoutLock() {
	for i := 0; i < 2000; i++ {
		x += 1
	}
}

func main() {
	x = 0
	for i := 0; i < 5; i++ {
		go addWithoutLock()
	}
	time.Sleep(time.Second)
	println("withoutLock: ", x)

	x = 0
	for i := 0; i < 5; i++ {
		go addWithLock()
	}
	time.Sleep(time.Second)
	println("withLock: ", x)
}
