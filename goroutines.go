package main

import (
	"fmt"
	"sync"
)
/*
	synchronizing all goroutines together
	this wg is saying that we are in group and will forward only when its done
*/
var wg = sync.WaitGroup{}
var counter int = 1;
var m = sync.RWMutex{}

func main () {
	/*
		A goroutine is a lightweight thread. It costs a lot less than a real OS thread, and multiple 
		goroutines may be multiplexed onto a single OS thread. The spec defines them as "an independent
		concurrent thread of control within the same address space".

		The go runtime is capable of handling thousands or even hundreds of thousands of goroutines 
		without a problem. As an example, the HTTP server in the standard lib handles all incoming
		requests by launching a new goroutine for each. Yet, it is capable of handling tens of thousands 
		q of requests per second for a typical request load (benchmark source).

		So all-in-all, just don't write "bad code". The goroutine scheduler is not (fully) preemptive, 
		so make sure your goroutines don't do senseless computation that would prevent the scheduler 
		to run other goroutines. Typically, system calls, IO operators and blocking operations
		(e.g. sending on / receiving from channels) are good yielding points. Many code do
		these under the hood even if you don't know it, so most code does not cause a problem.
		If one of your goroutines must do extensive calculations, you can always call 
		runtime.Gosched() to yield the processor, allowing other goroutines to run.
	*/
	/*
		A Mutex is used to provide a locking mechanism to ensure that only one Goroutine is running
	 	the critical section of code at any point in time to prevent race conditions from happening.
		runtime.GOMAXPROCS(number of threads that i want)
	*/

	for i := 0; i < 10; i++ {
		wg.Add(2);
		m.RLock()
		go sayHello();
		m.Lock()
		go increment()
		fmt.Printf("create #%v\n", i);
	}
	fmt.Println("here came for waiting")
	wg.Wait()
}
func sayHello(){
	fmt.Printf("Hello #%v\n", counter);
	m.RUnlock()
	wg.Done()
}

func increment(){
	counter++;
	m.Unlock()
	wg.Done()
}