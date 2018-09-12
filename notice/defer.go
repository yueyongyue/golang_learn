package main

import (
	"fmt"
)

func main() {
	DeferCall()
}

func DeferCall() {
	defer func() {fmt.Println("A")}()
	defer func() {fmt.Println("B")}()
	defer func() {fmt.Println("C")}()
	panic("error")
	fmt.Println("D")
}

/*
输出结果：
C
B
A
panic: error

defer 注册方式是保持在栈中，后入先出
panic 触发后defer依然执行，但panic后面的代码不在执行
*/
