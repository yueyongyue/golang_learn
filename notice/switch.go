package main

import "fmt"

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	Select("c")
	Select1("a")
}

func Select(args string) {
	switch args {
	case "a", "b":        // 只要匹配一个就执行
		fmt.Println("a b")
	case "c":
		fmt.Println("c")
		fallthrough       // 强制执行下面的case，无论后面的case是否为真
	case "d":
		fmt.Println("d")
	case "e":
		fmt.Println("e")
	default:
		fmt.Println("default")
	}
}

// switch的另一种写法
func Select1(args string) {
	switch {
	case args == "a":
		fmt.Println(args)
	default:
		fmt.Println("default")
	}
}