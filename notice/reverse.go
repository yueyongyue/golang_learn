package main

import "fmt"

func main() {
	s := "hello世界"
	fmt.Println(Reverse(s))
}

func Reverse(s string) string {
	// golang中的string类型存储的字符串是不可变的， 如果要修改string内容需要将string转换为[]byte或[]rune，并且修改后的string内容是重新分配的。
	// byte与rune类型有一个共性，即：它们都属于别名类型。byte是uint8的别名类型，而rune则是int32的别名类型
	// 有中文用rune
	r := []rune(s)
	for i, j := 0, len(r) - 1; i < j; i, j = i + 1, j - 1  {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}