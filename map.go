package main

import (
	"errors"
	"sync"
	"time"
)

// ConcurrentMap 是一个面向并发场景的 map
type ConcurrentMap struct {
	mu     sync.Mutex
	cond   *sync.Cond
	values map[string]interface{}
}

// NewConcurrentMap 创建一个新的 ConcurrentMap
func NewConcurrentMap() *ConcurrentMap {
	cm := &ConcurrentMap{
		values: make(map[string]interface{}),
	}
	cm.cond = sync.NewCond(&cm.mu)
	return cm
}

// Put 将 key-value 存入 map
func (cm *ConcurrentMap) Put(key string, value interface{}) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	cm.values[key] = value
	cm.cond.Broadcast() // 通知所有等待的goroutine
}

// Get 获取 key 对应的 value，如果 key 不存在则阻塞等待 maxWaitingTime 时间
func (cm *ConcurrentMap) Get(key string, maxWaitingTime time.Duration) (interface{}, error) {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	// 如果 key 存在，直接返回
	if val, ok := cm.values[key]; ok {
		return val, nil
	}

	// key 不存在，等待 maxWaitingTime 时间
	timeout := time.After(maxWaitingTime)
	for {
		cm.cond.Wait()

		// 检查 key 是否存在
		if val, ok := cm.values[key]; ok {
			return val, nil
		}

		// 检查是否超时
		select {
		case <-timeout:
			return nil, errors.New("timeout waiting for key")
		default:
			// 继续等待
		}
	}
}

func main() {
	// 创建一个 ConcurrentMap
	concurrentMap := NewConcurrentMap()

	// 启动一个goroutine往map中写入数据
	go func() {
		time.Sleep(2 * time.Second)
		concurrentMap.Put("exampleKey", "exampleValue")
	}()

	// 尝试从map中获取数据，设置最大等待时间为3秒
	result, err := concurrentMap.Get("exampleKey", 3*time.Second)
	if err != nil {
		println("Error:", err.Error())
	} else {
		println("Result:", result.(string))
	}
}
