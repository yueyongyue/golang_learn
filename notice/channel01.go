package main

import (
	"fmt"
	"time"
)

var c = make(chan int)

func ping() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	fmt.Println("Goroutine begin")
	c <- 0  // 向ch中加数据，如果没有其他goroutine来取走这个数据，那么挂起foo, 直到main函数把0这个数据拿走
	fmt.Println("Goroutine end")
}

func main() {
	go ping()
	// main函数被阻塞，从ch取数据，如果ch中还没放数据，那就挂起main线，直到foo函数中放数据为止
	time.Sleep(5*time.Second)
	fmt.Println(<- c)
	time.Sleep(5*time.Second)
	fmt.Println("Main end")
}

/*
输出：
Goroutine begin
Goroutine end
0
Main end

main()执行时读取c中数据，然后阻塞，直到c中有了数据
ping()的goroutine，执行时向c写入了数据


 */