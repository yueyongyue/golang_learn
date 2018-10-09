package main

import (
	"time"
	"fmt"
)

func main() {
	myChan := make(chan string)

	go func(){
		myChan <- "Message!"
	}()

	select {
	case msg := <- myChan:
		fmt.Println(msg)
	default:
		fmt.Println("No Msg")
	}

	<-time.After(time.Second * 1)

	select {
	case msg := <- myChan:
		fmt.Println(msg)
	default:
		fmt.Println("No Msg")
	}
}
