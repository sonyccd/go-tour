package main

import (
	"fmt"
	"time"
)

func ping(ch chan string) {
	for {
		ch <- "ping"
	}
}

func catcher(ch chan string) {
	for {
		msg := <-ch
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var ch chan string = make(chan string)

	go ping(ch)
	go catcher(ch)

	var input string
	fmt.Scanln(&input)
}
