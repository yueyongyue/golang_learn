package main

import "fmt"

type Goods struct {
	Price int
}

type Books struct {
	Goods *Goods // 使用指针的结构体组合
	Name string
	Author string
}
func main() {
	good := Goods{
		Price: 100,
	}
	fmt.Printf("goods的内存地址:%p\n", &good) // 0xc00001a108
	b := Books{
		Goods: &good,         // Goods 为指针类型，使用&取good的地址也就是0xc00001a108
		Name: "三国",
		Author:"jyp",
	}
	fmt.Printf("b的值:%+v\n", b) // {Goods:0xc00001a108 Name:三国}
	fmt.Printf("b的内存地址:%p\n", &b) //0xc00000a080
	fmt.Printf("b.Goods的值:%p\n", b.Goods) // 查看b.Goods的值，因为b.Goods本就是指针类型，可以直接用%p进行打印输出 0xc00001a108
	fmt.Printf("b.Goods的值指向的值%v\n", *b.Goods) // 查看b.Goods的值指向的真正的值， 使用*用于取指针类型的真正的值
	fmt.Printf("b.Goods的内存地址%v\n", &b.Goods) //0xc00006a150
	fmt.Printf("b.Name的内存地址%v\n", &b.Name) //0xc00006a158
	fmt.Printf("b.Author的内存地址%v\n", &b.Author) //0xc00006a168
}

// 1、在结构体组合结构体的时候，使用指针类型可以节省内存空间
// 2、结构体内存分配是连续的