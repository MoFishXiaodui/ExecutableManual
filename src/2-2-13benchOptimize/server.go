package bench

import "github.com/bytedance/gopkg/lang/fastrand"

var ServerIndex [10]int

// 初始化服务函数，（其实就是往数组里面各个项写入特定数字，i+100本身应该没什么含义）
func InitServerIndex() {
	for i := 0; i < 10; i++ {
		ServerIndex[i] = i + 100
	}
}

func Select() int {
	// 随意选择一个服务器（随机选择一个数组元素，模拟随机一个服务器返回数字）
	return ServerIndex[fastrand.Intn(10)]
}
