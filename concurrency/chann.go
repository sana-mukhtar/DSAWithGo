package main

import (
	"fmt"
	"time"
)

func printSomething(str string) {
	fmt.Println(str)
}

func main() {
	go printSomething("a")
	go printSomething("b")
	go printSomething("c")
	go printSomething("d")

	// channels -> channels are typicallly used to communicate information between go routines. channels are FIFO queues.main() is also a go routine.
	myChan := make(chan string)
	anotherChan := make(chan string)

	go func() {
		var data string
		time.Sleep(1 * time.Second)
		data = "dataaa"
		myChan <- data
	}()

	select {
	case messageFromMyChan := <-myChan:
		fmt.Println(messageFromMyChan)
	case messageFromAnotherChan := <-anotherChan:
		fmt.Println(messageFromAnotherChan)
	}

	// buffered channel
	buffChan := make(chan string)

	go func() {
		// time.Sleep(3*time.Second)
		buffChan <- "I am buffered chan1"
	}()

	go func() {
		buffChan <- "I am buffered chan2"
	}()
	msg := <-buffChan
	msg2 := <-buffChan
	fmt.Println("buff", msg)
	fmt.Println("buff2", msg2)

	// concurrency patterns
	// 1. for-select loop
	// 2. done channel
	// 3. pipelines

	// 1. for select loop
	chars := []string{"a", "b", "c"}
	newChan := make(chan string, 3)
	for _, ch := range chars {
		select {
		case newChan <- ch:
			fmt.Println(newChan)
		}
	}

	close(newChan)

	for chanVal := range newChan {
		fmt.Println(chanVal)
	}

	// 2. done channel
	myNewChan := make(chan string, 1)
	go func() {
		time.Sleep(time.Second)
		myNewChan <- "myNewChan value"
	}()

	select {
	case res := <-myNewChan:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("OOOPS timeout")
	}
}
