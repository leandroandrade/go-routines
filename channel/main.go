package main

import (
	"fmt"
	"time"
)

func main() {
	send := make(chan string)
	done := make(chan bool)

	go sendMessage(send)
	go receiveMessage(send, done)

	<-done

}

func sendMessage(send chan string) {
	fmt.Println("sending message...")

	time.Sleep(2 * time.Second)
	send <- "hello Gopher"

	fmt.Println("message sended!")
}

func receiveMessage(send chan string, done chan bool) {
	fmt.Println("waiting message...")

	fmt.Println("the message is", <-send)

	done <- true
	fmt.Println("message received!")
}
