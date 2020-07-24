package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Speeds interface {
	Ins()
	Loads() int64
}

// 普通的版本
type normal struct {
	counter int64
}

func (n normal) Ins () {
	n.counter++
}

func (n normal) Loads () int64 {
	return n.counter
}

// 锁版本
type LockNormal struct {
	counter int64
	lock sync.Mutex
}

func (l *LockNormal) Ins ()  {
	l.lock.Lock()
	defer l.lock.Unlock()

	l.counter++
}

func (l *LockNormal) Loads () int64 {
	l.lock.Lock()
	defer l.lock.Unlock()
	return l.counter
}

// 原子版本
type AtomicNormal struct {
	counter int64
}

func (a *AtomicNormal) Ins () {
	atomic.AddInt64(&a.counter, 1)
}

func (a *AtomicNormal) Loads () int64 {
	return atomic.LoadInt64(&a.counter)
}

func test (x Speeds) {
	var wg sync.WaitGroup
	start := time.Now()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			x.Ins()
			wg.Done()
		}()
	}
	wg.Wait()

	end := time.Now()
	fmt.Println(x.Loads(), end.Sub(start))
}


func main () {
	c1 := normal{}
	test(c1)
	c2 := LockNormal{}
	test(&c2)
	c3 := AtomicNormal{}
	test(&c3)
}
