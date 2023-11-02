package main

import (
	"fmt"
)

func isPrime(num int) bool {
	if num < 2 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var num int
	fmt.Print("请输入一个整数: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("输入错误:", err)
		return
	}

	if isPrime(num) {
		fmt.Printf("%d 是素数\n", num)
	} else {
		fmt.Printf("%d 不是素数\n", num)
	}
}
