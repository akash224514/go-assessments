package main

import "fmt"

// Function which is sending a sequential values over the channel
func myfun(mychnl chan int) {

	mychnl <- 10
	mychnl <- 20
	mychnl <- 35
	mychnl <- 100
	mychnl <- 200
	mychnl <- 502

	close(mychnl)
}

// Main function
func main() {
	fmt.Println("\nTitle - Create a small program with a for loop which can iterate over the sequential values sent on the channel until it closes")

	// Creating a channel
	c := make(chan int)

	// calling Goroutine
	go myfun(c)

	// When the value of ok is set to true means the
	// channel is open and it can send or receive data
	// When the value of ok is set to false means the channel is closed
	for {
		res, ok := <-c   // We are listening channel data via res variable
		if ok == false { //When data is over, we are  setting a flag to False
			fmt.Println("Channel Close ", ok)
			break
		}
		fmt.Println("Channel Open ", res, ok)
	}
}
