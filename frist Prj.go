package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomnumber := rand.Intn(101) + 0
	var low, high = 0, 100
	var mid = (low + high) / 2
	for {
		if randomnumber > mid {
			low = mid
			mid = (low + high) / 2
		}
		if randomnumber < mid {
			high = mid
			mid = (low + high) / 2
			if randomnumber == mid {
				fmt.Println(mid)
				break
			}
		}
	}
}
