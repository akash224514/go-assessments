package main

import (
	"fmt"
	"sync"
)

var x = 0
var myCh = make(chan int, 1) //Declaring a buffered channel named as myCh

func increment(wg *sync.WaitGroup) {
	myCh <- 1 // when one go routine try to access x, this buffered channel of capacity 1 will filled up and it will block other go routines from accessing x
	x = x + 1
	<-myCh // Reciver will unblock the buffered channel
	wg.Done()
}

func main() {

	fmt.Println("\n03 - Making use of Mutex to resolve the issue of Race condition")

	var w sync.WaitGroup

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
	}

	w.Wait()
	fmt.Println("final value of x", x)
	close(myCh)
}

/*
Approach to solve the problem ->

1) Here, initially for every call to an increment function, increment go routine is trying to access the vairable x and incrementing it by 1
2) Every time, for each i, one different increment go routine is trying to access x.
3) When we try to increment x by each call, there is a possibility of interleavings between those go routines.
4) As these go routiens are running concurrently and trying to access the variable x, It is producing different results each time.
5) Here, A buffered channel with size 1 has the property that if the single slot is filled, the channel will block the other go routines from accessing x.
6) Once the slot is vacated using a channel reciever, the channel will unblock and other go routine will now be able to access x likewise.
7) Likewise we are able to avoid interleaving of go routines with the help of channel and thus producing a result appropriately each time.

*/
