package main

import "fmt"

type Goods struct {
	Price int
}
type Vendor struct {
	Name string
}

type Books struct {
	Goods  *Goods // 使用指针的结构体组合
	Name   string
	Author string
	Date   *int
	Vendor Vendor
}

func main() {
	good := Goods{
		Price: 100,
	}
	fmt.Printf("goods的内存地址:%p\n", &good) // 0xc00001a108
	b := Books{
		Goods:  &good, // Goods 为指针类型，使用&取good的地址也就是0xc00001a108
		Name:   "三国演义",
		Author: "jyp",
		Vendor: Vendor{Name:"邮电出版社"},
	}
	fmt.Printf("b的内存地址:%p\n", &b)            //0xc00000a080
	fmt.Printf("b.Goods的值:%p\n", b.Goods)    // 查看b.Goods的值，因为b.Goods本就是指针类型，可以直接用%p进行打印输出 0xc00001a108
	fmt.Printf("b.Goods的内存地址%v\n", &b.Goods) //0xc00006a150
	b.Modify()
	fmt.Println("b.Goods.Price的值:", b.Goods.Price)
	fmt.Println("b.Vendor.name的值:", b.Vendor.Name)
	b.Modify1()
	fmt.Println("b.Goods.Price的值:", b.Goods.Price)
	fmt.Println("b.Vendor.name的值:", b.Vendor.Name)
}

func (b Books) Modify() {
	fmt.Println("#######################")
	fmt.Printf("b的内存地址:%p\n", &b)            // 方法传递的不是指针，内存地址变更，进行值拷贝
	fmt.Printf("b的内存地址:%+v\n", b)            // 方法传递的不是指针，内存地址变更，进行值拷贝
	fmt.Printf("b.Goods赋值前的值:%p\n", b.Goods) // 为值拷贝，b.Goods的值(也就是一个内存地址)也不会变
	// b.Goods的值变成新的值了， 原始的b.Goods不会改变
	//b.Goods = &Goods{
	//	Price: 999,
	//}
	// 非指针传递的结构体方法，想要改变之前为指针的字段，可以使用如下方法， b.Goods存储的指针值没有变，想要进行修改可以使用*(b.Goods)或者
	// *b.Goods(先进行.计算，在镜像*运算,带上括号更容易理解)获取指针值指向的真正值，进行赋值修改，原始的b.Goods也会改变
	//*(b.Goods) = Goods{
	//	Price: 66666,
	//}
	// 原始的b.Goods会改变
	b.Goods.Price = 999999
	// 原始的b.Goods会改变 *(b.Goods)取值Goods对应的真正的值，在为price赋值
	(*(b.Goods)).Price = 11111
	// 原始的b.Vendor.Name会改变
	b.Vendor.Name = "中信出版社"
	fmt.Printf("b.Goods赋值后的值:%p\n", b.Goods)
	fmt.Printf("b修改后的内存地址:%p\n", &b)
}

func (b *Books) Modify1() {
	fmt.Println("****************************")
	fmt.Printf("b的内存地址:%p\n", b)            // 方法传递的不是指针，内存地址变更，进行值拷贝
	fmt.Printf("b.Goods赋值前的值:%p\n", b.Goods) // 为值拷贝，b.Goods的值(也就是一个内存地址)也不会变
	// 原始的b.Goods会改变
	//b.Goods = &Goods{
	//	Price: 88888,
	//}
	// 原始的b.Goods会改变 Go提供的隐式解引用特性
	//b.Goods.Price = 33333
	// 原始的b.Goods会改变
	(*((*b).Goods)).Price = 565656
	// 应该这么写，编译器直接简化了，如上方式即可重新赋值
	// (*b)取得对应的真正的值，(*b).Goods 获取Goods字段的值(值为内存地址)，*((*b).Goods)获取真正的值，在进行赋值
	//*((*b).Goods) = Goods{
	//	Price: 77777,
	//}
	// 原始的b.Vendor.Name会改变
	//b.Vendor.Name = "机械出版社"
	// 原始的b.Vendor.Name会改变
	(*b).Vendor.Name ="人民邮电出版社"
	fmt.Printf("b.Goods赋值后的值:%p\n", b.Goods)
}

// 只要是指针类型的结构体方法，无论字段的类型是否为指针，在方法中进行修改都会改变原有的值
// 非指针类型的结构体方法，字段为指针类型的，在方法中进行修改，会修改原来的值，字段为非指针类型的则不能修改