package main

import (
	"os"
	"os/signal"
	"fmt"
)

func main() {
	fmt.Println(os.Getpid())
	c := make(chan os.Signal, 0)
	signal.Notify(c)
	s := <- c
	fmt.Println("Got signal:", s)
}
