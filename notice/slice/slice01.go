package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a = [5]int{1, 2, 3, 4, 5}
	s := a[1:4]
	fmt.Printf("数组a的值:%v\n", a)
	fmt.Printf("数组a的内存地址:%p\n", &a)
	fmt.Printf("数组a[1]的内存地址:%p\n", &a[1])
	fmt.Printf("切片s的值:%v\n", s)
	fmt.Printf("切片s的内存地址:%p\n", &s)
	fmt.Printf("切片s的真正的值:%p\n", s)
	var ptr unsafe.Pointer
	var length = 3
	var s1 = struct {
		addr unsafe.Pointer
		len  int
		cap  int
	}{ptr, length, length}
	fmt.Println(s1)
	s2 := *(*[]byte)(unsafe.Pointer(&s1))
	fmt.Println(s2)
	var a1 = [6]int{10, 20, 30, 40, 50, 60}
	s3 := a1[2:5:6]
	fmt.Println(s3)
	fmt.Println(len(s3))
	fmt.Println(cap(s3)) // a[low  high : max]  容量由max - low 定义

	newSlice := new([100]int)[0:50]    //这个两个是等效的
	makeSlice := make([]int, 50, 100)
	fmt.Println(newSlice)
	fmt.Println(makeSlice)

}
