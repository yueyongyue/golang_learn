package main

import (
	"fmt"
	"strings"
)

/*
Codewars.accum("abcd") // -> "A-Bb-Ccc-Dddd"
Codewars.accum("RqaEzty") // -> "R-Qq-Aaa-Eeee-Zzzzz-Tttttt-Yyyyyyy"
Codewars.accum("cwAt") // -> "C-Ww-Aaa-Tttt"
 */

 func main() {
 	fmt.Println(Accum("cwAt"))
 	fmt.Println(Accum1("cwAt"))
 	fmt.Println([]byte("a"))
 	fmt.Println([]byte("z"))
 	fmt.Println([]byte("A"))
 	fmt.Println([]byte("Z"))
 }

 func Accum(s string) string {
 	var tmp string
 	for i, v := range s {
		var lower string
 		for j :=0; j < i; j ++ {
 			lower += strings.ToLower(string(v))
		}
        tmp += strings.ToUpper(string(v)) + lower + "-"
	}
	return strings.TrimRight(tmp, "-")
 }

 func Accum1(s string) string {
 	l := make([]string, len(s))
 	for i, v := range s {
 		l[i] = strings.ToUpper(string(v)) + strings.Repeat(strings.ToLower(string(v)), i)
	}
	return strings.Join(l, "-")
 }