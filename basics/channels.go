package main

import (
	"fmt"
	"sync"
)
/*
* channels are biderection and alloew to send the data of specific type from that

flow is like this
1. sender first check that (i)if there is any reciever is there if yes send the data into it (ii) if
not wait till some receiver is there
2. receiver waits untill something is come

* by default channels are biderection but we can also send the send-only or receive-only kind of 
channel to go routines

* buffered channels ---> store some buffer is channel
used when receiver freq < senders freq

* we can close the channels to let know on receiver that nothing is coming more

* we should always have the strategy that how we are going to shut down the go routine

* we use informerChannel = make(chan struct{})//this do not consume any memory

* we can use it with select

*/
var wg = sync.WaitGroup{}

func recOnly(ch <-chan int){ //<-chan says can only exctract something from channel
	// for i := range ch {
	// 	fmt.Println(i)
	// }
	for {
		i, ok := <- ch;
		if ok {
			fmt.Println(i)
		} else {
			break
		}
	}
	wg.Done()
}

func sendOnly(ch chan<- int){ //chan<- says can only exctract something from channel
	for i := 1; i < 10; i++ {
		ch <- i;
	}
	close(ch)
	wg.Done()
}

func main(){
	//ch := make(chan int);
	ch := make(chan int, 50);//buffered channel where size is 50 integers
	wg.Add(2);
	go recOnly(ch)
	go sendOnly(ch)
	wg.Wait()
}