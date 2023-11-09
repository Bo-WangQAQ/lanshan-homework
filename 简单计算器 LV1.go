package main

import "fmt"

func main() {
	var a, b int
	var ch string
	fmt.Println("请输入算式：")
	fmt.Scan(&a, &ch, &b)
	switch ch {
	case "+":
		fmt.Println(a, ch, b, "=", a+b)
	case "-":
		fmt.Println(a, ch, b, "=", a-b)
	case "*":
		fmt.Println(a, ch, b, "=", a*b)
	case "/":
		fmt.Println(a, ch, b, "=", a/b)
	default:
		fmt.Println("输入有误！")
	}

}
