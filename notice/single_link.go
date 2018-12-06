package main

import "fmt"

type HeroNode struct {
	No       int
	Name     string
	NickName string
	Next     *HeroNode
}

// 单链表插入
func InsertHeroNode(head, newHero *HeroNode) {
	temp := head
	for {
		if temp.Next == nil { // 链表Next为nil的时候说明到达链表最后
			temp.Next = newHero // 在最后插入数据
			break
		}
		temp = temp.Next
	}
}

// 显示链表
func ListHeroNode(head *HeroNode) {
	temp := head
	if temp.Next == nil {
		fmt.Println("空链表")
		return
	}
	for {
		fmt.Printf("[%d, %s, %s]==>", temp.Next.No, temp.Next.Name, temp.Next.NickName)
		temp = temp.Next
		if temp.Next == nil {
			break
		}
	}
}
func main() {
	head := &HeroNode{}

	hero1 := &HeroNode{
		No:       1,
		Name:     "宋江",
		NickName: "及时雨",
	}
	hero2 := &HeroNode{
		No:       2,
		Name:     "卢俊义",
		NickName: "玉麒麟",
	}
	hero3 := &HeroNode{
		No:       3,
		Name:     "林冲",
		NickName: "豹子头",
	}
	InsertHeroNode(head, hero1)
	InsertHeroNode(head, hero2)
	InsertHeroNode(head, hero3)
	ListHeroNode(head)
}
