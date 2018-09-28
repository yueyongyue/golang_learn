package main

import "fmt"
// Golang中只有三种引用类型：slice(切片)、map(字典)、channel(管道)
func main() {
	a := 10
	fmt.Println(&a)
	// 值类型 , 值拷贝
	b := a
	fmt.Println(&b) // 两次的内存地址不同

	l := []int{1,2,3,4,5}
	z := l
	fmt.Printf("%p %v\n", l,l)
	fmt.Printf("%p %v\n", z,z)
	z[0] = 100     // 数组l也会改变 l和z本质上指向同一个底层数组
	fmt.Println(l)
	fmt.Println(z)
}
