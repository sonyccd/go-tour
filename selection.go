package main

import (
	"fmt"
	"time"
)

func main() {
	//bidirectional chan
	ch1 := make(chan string)
	ch2 := make(chan string)
	//anon function
	go func() {
		for {
			ch1 <- "from ch1"
			time.Sleep(time.Second * 1)
		}
	}() // have to add the () at the end to invoke the function

	go func() {
		for {
			ch2 <- "from ch2"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			//which one is avalible do that
			select {
			case msg1 := <-ch1:
				fmt.Println(msg1)
			case msg2 := <-ch2:
				fmt.Println(msg2)
			default:
				fmt.Print(".")
				//if you put time.Sleep(time.Second * 3)
				//it will block the entire select for that time
				//if there is not buffer than msgs will be missed
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
