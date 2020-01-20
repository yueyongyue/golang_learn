package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var l []int
	var m sync.Mutex
	for i:= 0; i < 1000; i++ {
		go func(i int) {
			m.Lock()
			l = append(l, i)
			m.Unlock()
		}(i)
	}
	var a int
	var mm sync.Mutex
	for i:= 0; i < 1000; i++ {
		go func(i int) {
			mm.Lock()
			a += 1
			mm.Unlock()
		}(i)
	}
	time.Sleep(time.Second)
    fmt.Println(len(l))
	fmt.Println(a)
}

/*
不使用互斥锁最后的值不正确
协程依次执行：从寄存器读取 a 的值 -> 然后做加法运算 -> 最后写到寄存器。试想，此时一个协程取出 a 的值 3，正在做加法运算（还未写回寄存器）。同时另一个协程此时去取，取出了同样的 a 的值 3。最终导致的结果是，两个协程产出的结果相同，a 相当于只增加了 1。

所以，锁的概念就是，我正在处理 a（锁定），你们谁都别和我抢，等我处理完了（解锁），你们再处理。这样就实现了，同时处理 a 的协程只有一个，就实现了同步。
 */