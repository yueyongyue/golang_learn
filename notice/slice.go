package main

import "fmt"

func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
	s1 := []int{}
	s1 = append(s1, 1, 2, 3)
	fmt.Println(s1)
}

/*
输出：
[0 0 0 0 0 1 2 3]
[1 2 3]

make([]int, 5) 之后，这个list就变成[0, 0, 0, 0, 0] 了，默认0填充。
*/
