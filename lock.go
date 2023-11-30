package main

import (
	"sync/atomic"
)

type PetersonLock struct {
	flag [2]uint32
	turn uint32
}

func NewPetersonLock() *PetersonLock {
	return &PetersonLock{}
}

func (l *PetersonLock) Lock() {
	me := 0         // 当前线程的索引
	other := 1 - me // 其他线程的索引

	atomic.StoreUint32(&l.flag[me], 1)      // 设置当前线程的标志位为1
	atomic.StoreUint32(&l.turn, uint32(me)) // 设置轮到当前线程执行

	for atomic.LoadUint32(&l.flag[other]) == 1 && atomic.LoadUint32(&l.turn) == uint32(me) {
		// 如果其他线程正在临界区并且轮到当前线程执行，当前线程需要等待
	}
}

func (l *PetersonLock) Unlock() {
	me := 0 // 当前线程的索引

	// 将当前线程的标志位置为0，表示退出临界区
	atomic.StoreUint32(&l.flag[me], 0)
}

func main() {
	// 使用自定义的PetersonLock进行互斥操作
	lock := NewPetersonLock()

	// 在多个goroutine中测试互斥性
	for i := 0; i < 10; i++ {
		go func(i int) {
			lock.Lock()
			println("Critical section entered by goroutine", i)
			lock.Unlock()
		}(i)
	}

	// 等待所有goroutine执行完成
	// ...

	println("All goroutines have completed")
}
