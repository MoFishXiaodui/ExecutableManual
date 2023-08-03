package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFirstLine() string {
	file, err := os.Open("log.txt")
	defer file.Close()
	if err != nil {
		return ""
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		return scanner.Text()
	}
	return ""
}

func ProcessFirstLine() string {
	line := ReadFirstLine()
	destLine := strings.ReplaceAll(line, "11", "00")
	return destLine
}

func main() {
	line := ProcessFirstLine()
	fmt.Println(line)
}
