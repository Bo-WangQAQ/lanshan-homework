package main

import (
	"fmt"
	"sync"
)

// printNumbers 是一个打印数字的函数，根据 isOdd 决定打印奇数或偶数
func printNumbers(wg *sync.WaitGroup, ch1, ch2 chan int) {
	defer wg.Done()

	for i := 1; i <= 100; i += 2 {
		// 向 ch1 发送奇数
		ch1 <- i
		// 从 ch2 接收上一个 goroutine 发送的值
		val := <-ch2

		// 打印奇数
		fmt.Printf("A: %d\n", <-ch1)
		// 打印偶数
		fmt.Printf("B: %d\n", val)
	}
}

func main() {
	var wg sync.WaitGroup
	// 创建两个带缓冲的通道，缓冲区大小为1，用于 goroutine 之间的通信
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	wg.Add(2)

	// 启动第一个 goroutine
	go func() {
		defer wg.Done()
		// 第一个 goroutine 以奇数开始
		printNumbers(&wg, ch1, ch2)
	}()

	// 启动第二个 goroutine
	go func() {
		defer wg.Done()
		// 第二个 goroutine 以偶数开始
		printNumbers(&wg, ch1, ch2)
	}()

	// 初始启动，通过 ch2 向第一个 goroutine 发送初始值
	ch2 <- 2

	// 等待两个 goroutine 完成
	wg.Wait()

	// 关闭通道
	close(ch1)
	close(ch2)
	fmt.Println("Program finished.")
}
