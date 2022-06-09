package main

import (
	"fmt"
	"sync"
)

var x = 0
var mut sync.Mutex //Declaring a mutex variable using sync package

func increment(wg *sync.WaitGroup) {
	mut.Lock() // Adding a mutex lock in order to avoid race condition
	x = x + 1
	mut.Unlock() // Adding mutex unlock so that another go routine can access the operation
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
}

/*
Approach to solve the problem ->

1) Here, initially for every call to an increment function, increment go routine is trying to access the vairable x and incrementing it by 1
2) Every time, for each i, one different increment go routine is trying to access x.
3) When we try to increment x by each call, there is a possibility of interleavings between those go routines.
4) As these go routiens are running concurrently and trying to access the variable x, It is producing different results each time.
5) when we add lock (lock is mutex method), It is siganling other go routines that variable x is in use. and it increments x by 1
6) After incrementing x by 1,we are unlocking the mutex so that other go routine can access x and increment it by 1
7) Likewise we are able to avoid interleavings of go routines with the help of mutex and thus producing a result appropriately each time.

*/
