package main

import "fmt"

func main() {
	i := GetValue()
	switch i.(type) {
	case int: {
		fmt.Println("int")
	}
	case string: {
		fmt.Println("string")
	}
	case interface{}: {
		fmt.Println("interface")
	}
	default: {
		fmt.Println("unknown")
	}
	}
}

func GetValue() int {
	return 1
}

func GetValue1() interface{} {
	return 1
}

/*
输出：

.\interview03.go:7:2: cannot type switch on non-interface value i (type int)
type只能用在interface上，使用GetValue1()即可
*/
