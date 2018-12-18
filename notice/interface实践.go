package main

import (
	"sort"
	"fmt"
	"math/rand"
)

type Hero struct {
	Name string
	Age int
}

// 声明一个Hero结构体的切片类型
type HeroSlice []Hero

// 实现Interface接口
func (h HeroSlice) Len() int {
	return len(h)
}

// Less方法就是决定你使用什么标准进行排序
// 按照Hero的年龄从小到大排序
func (h HeroSlice) Less(i, j int) bool {
	//return  h[i].Age < h[j].Age
	return  h[i].Name < h[j].Name
}

func (h HeroSlice) Swap(i, j int) {
	//temp := h[i]
	//h[i] = h[j]
	//h[j] = temp
	h[i], h[j] = h[j], h[i]
}

func main() {
	intSlice := []int{2, -2, 0, 9, 4}
	sort.Ints(intSlice)
	fmt.Println(intSlice)
	// 对结构体进行排序
	// 使用 sort包里面的Sort方法
	heroSlice := HeroSlice{}
	for i := 0; i < 10; i ++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄~%d", rand.Intn(100)),
			Age: rand.Intn(100),
		}
		heroSlice = append(heroSlice, hero)
	}
	for _, v := range heroSlice {
		fmt.Println(v)
	}
	fmt.Println()
	// 只要 heroSlice的类型实行了sort.Sort中Interface这个接口，就可以使用sort.Sort进行排序
	sort.Sort(heroSlice)
	for _, v := range heroSlice {
		fmt.Println(v)
	}
}
