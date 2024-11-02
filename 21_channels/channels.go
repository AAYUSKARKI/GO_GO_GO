package main

import (
	"fmt"
)

// func processNum(numChan chan int) {
// 	fmt.Println("Process number:", <-numChan)
// }

// func task(done chan bool) {
// 	defer func() {
// 		done <- true
// 	}()
// 	fmt.Println("Task started")
// }

// func emailSender(emailChan chan string, done chan bool) {

// 	defer func() {
// 		done <- true
// 	}()

//		for email := range emailChan {
//			fmt.Println("Sending email to:", email)
//			time.Sleep(time.Second)
//		}
//	}
func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 42
	}()

	go func() {
		chan2 <- "Hello, World!"
	}()

	for i := 0; i < 2; i++ {
		select {
		case num := <-chan1:
			fmt.Println("Received number from channel 1:", num)
		case msg := <-chan2:
			fmt.Println("Received message from channel 2:", msg)
		}
	}
	// emailChan := make(chan string)
	// done := make(chan bool)

	// go emailSender(emailChan, done)

	// for i := 0; i < 100; i++ {
	// 	emailChan <- fmt.Sprintf("email-%d", i)
	// }

	// close(emailChan)

	// <-done // blocks until task is done

	// numChan := make(chan int)

	// go processNum(numChan)

	// numChan <- 42

	// time.Sleep(1 * time.Second)
	// messageChan := make(chan string)

	// messageChan <- "Hello, World!"

	// msg := <-messageChan

	// fmt.Println(msg)
}
