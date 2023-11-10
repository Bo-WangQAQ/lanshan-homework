package main

import (
	"fmt"
)

func add(x, y float64) float64 {
	return x + y
}

func subtract(x, y float64) float64 {
	return x - y
}

func multiply(x, y float64) float64 {
	return x * y
}

func divide(x, y float64) float64 {
	if y != 0 {
		return x / y
	}
	panic("division by zero")
}

func main() {
	var operator string
	var num1, num2 float64

	fmt.Print("请输入第一个数字: ")
	fmt.Scanln(&num1)

	fmt.Print("请输入第二个数字: ")
	fmt.Scanln(&num2)

	fmt.Print("请输入操作符(+, -, *, /): ")
	fmt.Scanln(&operator)

	var result float64

	switch operator {
	case "+":
		result = add(num1, num2)
	case "-":
		result = subtract(num1, num2)
	case "*":
		result = multiply(num1, num2)
	case "/":
		result = divide(num1, num2)
	default:
		fmt.Println("无效的操作符")
		return
	}

	fmt.Printf("结果: %.2f\n", result)
}
