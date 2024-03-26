package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in GOlang")

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}


	// myCh <- 5  
	// fmt.Println(<-myCh)  // 5
	wg.Add(2)
	go func (ch chan int, wg *sync.WaitGroup)  {
// fmt.Println(<-myCh)
close(myCh)
val, isChannelOpen := <- myCh
fmt.Println(isChannelOpen)
fmt.Println(val)
		wg.Done()
		
	}(myCh, wg)
	go func( ch chan int, wg *sync.WaitGroup){
		
		myCh <- 0
		close(myCh)
		myCh <- 6
		
		wg.Done()
	}(myCh,wg)

}

