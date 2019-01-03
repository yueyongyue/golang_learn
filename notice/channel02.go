package main

import "fmt"

var ch = make(chan int, 20000)
var tmp = make(chan int, 8)
var resChan = make(chan map[string]int, 20000)

func Write() {
	for i := 0; i < 20000; i ++ {
		ch <- i
	}
	close(ch)
}

func Read() {
	for v := range ch {
		var res int
		for i:= 1; i <= v;i ++ {
			res += i
		}
		resChan <- map[string]int{fmt.Sprintf("res[%d]",v):res}
	}
	tmp <- 0
}

func main() {
	go Write()
	for i := 0; i < 8; i ++ {
		go Read()
	}
	for {
		if len(tmp) == 8 {
			close(tmp)
			close(resChan)
			break
		}
	}
	for r := range resChan {
		for k, v := range r {
			fmt.Printf("%v=%d\n", k,v )
		}
	}
	fmt.Println("完成")
}
