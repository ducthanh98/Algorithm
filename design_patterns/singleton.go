package main

import (
	"fmt"
	"sync"
	"time"
)

var counter ISingleton
var locker sync.Once

type ISingleton interface {
	Incr() int
}

type Singleton struct {
	count int
}

func GetInstance() ISingleton {
	//locker.Lock()
	//if counter == nil {
	//	counter = &Singleton{count: 100}
	//}
	//locker.Unlock()
	locker.Do(func() {
		if counter == nil {
			counter = &Singleton{count: 100}
		}
	})

	return counter
}

func (s *Singleton) Incr() int {
	s.count++
	return s.count
}

func main() {
	for i := 0; i < 100; i++ {
		go func() {
			tmp := GetInstance()
			fmt.Printf("%p\n", tmp)
		}()
	}
	time.Sleep(5 * time.Second)
}
