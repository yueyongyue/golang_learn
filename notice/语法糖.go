package main

import "fmt"

// ... 的使用
func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	str1 := []string{
		"a",
		"b",
		"c",
	}
	str2 := []string{
		"x",
		"y",
		"z",
	}
	str1 = append(str1, str2...)
	fmt.Println(str1)
	Test(str2...) // 切片被打散
}

func Test(args ...string) {
	fmt.Println(args)
}