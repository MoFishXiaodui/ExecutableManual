package coverage

func Judge(n int) bool {
	if n >= 60 {
		return true
	}

	// 观察通过率时加一些只增加行数，不改变逻辑的代码
	// n += 10
	// n -= 10

	return false
}
