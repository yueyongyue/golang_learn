package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func foo() {
	fmt.Println("Goroutine begin")
	ch <- 0  // 向ch中加数据，如果没有其他goroutine来取走这个数据，那么挂起foo, 直到main函数把0这个数据拿走
	fmt.Println("Goroutine end")
}

func main() {
	go foo()
	// main函数被阻塞，从ch取数据，如果ch中还没放数据，那就挂起main线，直到foo函数中放数据为止
	time.Sleep(5*time.Second)
	fmt.Println("Main end")
}

/*
输出：
Goroutine begin
Main end

foo()的goroutine，执行时向ch写入了数据，如果没有其他的goroutine取出ch的数据，foo()的goroutine就阻塞了，所以只打印了开始而没有结束，
此goroutine会随着主程序的结束而结束，最后打印“Main end”

 */