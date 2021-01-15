package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var lock sync.Mutex
var rwlock sync.RWMutex
var num int

func main() {
	// WaitGroup的使用
	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	for i := 0; i < 10000; i++ {
	//		fmt.Println(i)
	//	}
	//}()
	//defer wg.Wait()
	//fmt.Println("hello")

	// lock的使用
	//wg.Add(1)
	//go add()
	//
	//wg.Add(1)
	//go sub()
	//
	//wg.Wait()
	//fmt.Println(num)

	// relock的使用
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go read()
	}
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go write()
	}
	wg.Wait()
}

func add() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		lock.Lock()
		num += 1
		lock.Unlock()
	}
}

func sub() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		lock.Lock()
		num -= 1
		lock.Unlock()
	}
}

func read() {
	defer wg.Done()
	rwlock.RLock() // 上读写锁, 读写锁之间不会相互阻塞, 只会和普通锁相互阻塞
	fmt.Println("查询数据")
	time.Sleep(time.Second * 5)
	fmt.Println("查询ok")
	rwlock.RUnlock()
}

func write() {
	defer wg.Done()
	rwlock.Lock()
	fmt.Println("写入数据")
	time.Sleep(time.Second * 5)
	fmt.Println("写入ok")
	rwlock.Unlock()
}
