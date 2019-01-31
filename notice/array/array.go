package main

import "fmt"

func main() {
	var a [3]int
	b := [3]int{}
	var c = [...]int{}
	var d = [3]int{1, 2, 3}
	var e = [3]int{1: 10, 0: 11, 2: 12}
	fmt.Println(a, b, c, d, e)
	fmt.Printf("d的内存地址%p\nd[0]的内存地址%p\nd[1]的内存地址%p\nd[2]的内存地址%p\n",
		&d, &d[0], &d[1], &d[2])
	/*
	d的内存地址0xc000014500         数组的内存地址就是数组首个元素的地址
	d[0]的内存地址0xc000014500
	d[1]的内存地址0xc000014508      相差8
	d[2]的内存地址0xc000014510      相差8
	 */
	Change(a) // 值类型不会被更改
	fmt.Println("a:", a)
	Change1(&a) // 传递指针，值被改变
	fmt.Println("a:", a)
	var z [26]byte
	z[0] = 'A' + 1  //双引号是用来表示字符串string，其实质是一个byte类型的数组，单引号表示rune类型。还有一个反引号，用来创建原生的字符串字面量
	fmt.Println(z)
	fmt.Println(string(z[0]))
}

func Change(a [3]int) {
	a[0] = 100
}

func Change1(a *[3]int) {
	a[0] = 1000
	//(*a)[0] = 1000
}

/*
1、数组在内存中的分配是连续的
2、数组的内存地址就是数组首个元素的地址
3、一个数组一旦声明/定义，其大小是固定的，不能动态增长
4、数组是值类型，默认是值传递，因此会进行值拷贝，数据间不会相互影响
5、[3]int与[10]int不是同一个数据类型 [...]int{1,2,3,4} 的类型为[4]int, 根据声明的
6、var z [...]int  这样声明会报错 use of [...] array outside of array literal
 */
