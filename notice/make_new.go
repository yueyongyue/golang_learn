package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("make slice map err: %s", err)
		}
	}()

	b := map[string]string{}
	fmt.Println(b)
	b["B"] = "B"
	fmt.Println(b)

	// make用于内建类型（只能用于创建map、slice 和channel）的内存分配。并且返回一个有初始值(非零)的T类型，而不是*T
	c := make(map[string]string)
	fmt.Println(c)
	c["C"] = "C"
	fmt.Println(c)

	// 返回了一个指针，指向新分配的类型T的零值。有一点非常重要：*new返回指针。
	d := new([]int)
	fmt.Println(d)

	var a map[string]string
	fmt.Println(a)
	a["A"] = "A"
	fmt.Println(a)
}
