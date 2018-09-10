package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	done := make(chan bool)

	go task(done)
	wg.Add(1)

	go loader(done)

	wg.Wait()

}

// just write chan, (<- done) not work
func task(done chan<- bool) {
	fmt.Println("start task")
	time.Sleep(3 * time.Second)

	done <- true
}

// just read from chan
func loader(done <-chan bool) {
	defer wg.Done()

	var i int
	chars := []rune(`|\- /`)

	for {
		select {
		case <-done:
			fmt.Printf("\r")
			fmt.Println("done!")
			return
		default:
			fmt.Printf("\r")
			fmt.Printf(string(chars[i]))

			i++

			if i == len(chars) {
				i = 0
			}

		}
	}
}
