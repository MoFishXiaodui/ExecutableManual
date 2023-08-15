package main

import (
	"fmt"
	"web/repository"
)

func main() {
	err := repository.InitTopicIndexMap("../data/")
	fmt.Println("err", err)
}
