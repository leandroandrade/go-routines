package main

import (
	"fmt"
	"time"
)

func every500ms(c1 chan string) {
	for {
		c1 <- "Every 500ms"
		time.Sleep(time.Millisecond * 500)
	}
}

func everyTwoSenconds(c2 chan string) {
	for {
		c2 <- "Every two seconds"
		time.Sleep(time.Second * 2)
	}
}

func processDataChannels(c1, c2 chan string) {
	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go every500ms(c1)
	go everyTwoSenconds(c2)

	processDataChannels(c1, c2)
}
