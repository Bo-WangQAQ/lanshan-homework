package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("请输入两个需要相加的数：")
	fmt.Scan(&a, &b)
	fmt.Println("两数之和为：", a+b)
	return
}
