package main

import (
	"strings"
	"fmt"
)

func main() {
	fmt.Println(band("knife"))
	fmt.Println(band("sandles"))
	fmt.Println(band("The Step-daughter"))
	fmt.Println(A("Cat Step,daughter"))
	fmt.Println(A(strings.Title("cat step,daughter")))
	fmt.Println(A(strings.Title("sandlesandles")))
	m := strings.Map(func(r rune) rune {
		return r - 32
	},"abc")
	fmt.Println(m)
}

func band(s string) string {
	b := []byte(A(s))
	if b[0] != b[len(b)-1] {
		if string(b[0:3]) == "The" {
			return strings.ToUpper(string(b[0])) + string(b[1:])
		} else {
			return "The " + strings.ToUpper(string(b[0])) + string(b[1:])
		}
	} else {
		return  strings.ToUpper(string(b[0])) + string(b[1:len(b)-1]) + string(b)
	}
}

func A(s string) string {
	str := ""
	b := []rune(s)
	for i := 0; i < len(b); i ++ {
		if b[i] < 65 || b[i] > 122 || b[i] > 90 && b[i] < 97 {
			if b[i+1] >= 65 && b[i+1] <= 90 {
				str += string(b[i]) + string(b[i+1])
			} else {
				str += string(b[i]) + string(b[i+1]-32)
			}
			i ++
		} else {
			str += string(b[i])
		}
	}
	return str
}