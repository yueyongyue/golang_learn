package main

import (
	"fmt"
	"sync"
)
/*
go提供了sync包和channel来解决协程同步和通讯问题
sync.WaitGroup只有3个方法，Add()，Done()，Wait()。
其中Done()是Add(-1)的别名。简单的来说，使用Add()添加计数，Done()减掉一个计数，计数不为0, 阻塞Wait()的运行。
 */
func main () {
	wg := sync.WaitGroup{}
	for i := 1; i < 4; i ++ {
		wg.Add(i)           //一般这么写 wg.Add(1),  就可以用 wg.Done()减掉计数
	}
	go func() {
		fmt.Println("Goroutine A")
		wg.Done()           // wg.Add(-1)
	}()
	go func() {
		fmt.Println("Goroutine B")
		wg.Add(-2)    // 开始的时候加了2这里加-2即减掉了这个计数
	}()
	go func() {
		fmt.Println("Goroutine C")
		wg.Add(-3)
	}()
	wg.Wait()
	fmt.Println("Main Process")
}