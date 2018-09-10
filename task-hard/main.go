package main

import (
	"fmt"
	"time"
	"sync"
)

var wg sync.WaitGroup

func task(name string) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		time.Sleep(1 + time.Second)

		fmt.Printf("hard task %v...\n", name)
	}
	fmt.Printf("task %v done!\n", name)
}

func otherTask() {
	time.Sleep(1 + time.Second)
	fmt.Println("========== EASYTASK ==========")
}

func main() {
	for i := 1; i <= 10; i++ {
		wg.Add(1)

		go task(fmt.Sprintf("%v", i))
	}

	otherTask()

	wg.Wait()
	fmt.Println("done!")
}
