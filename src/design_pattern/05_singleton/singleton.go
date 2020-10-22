package _5_singleton

import "sync"

type Singleton struct {
}

var (
	singleton *Singleton
	once      sync.Once
)

func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})
	return singleton
}
