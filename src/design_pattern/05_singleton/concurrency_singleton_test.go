package _5_singleton

import (
	"sync"
	"testing"
)

const pCount = 100

// 100并发协程获取单例
func TestConcurrencySingleton(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(pCount)
	instances := [pCount]*Singleton{}
	for i := 0; i < pCount; i++ {
		go func(index int) {
			instances[index] = GetInstance()
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 1; i < pCount; i++ {
		if instances[i] != instances[i-1] {
			t.Fatal("instance-", i, "not equal instance-", i-1)
		}
	}
}
