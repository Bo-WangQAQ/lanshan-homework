package main

import (
	"fmt"
	"math"
)

func main() {
	var r float32
	fmt.Println("请输入圆的半径：")
	fmt.Scan(&r)
	fmt.Println("所求圆的面积为：", r*r*math.Pi/2)
	return
}
