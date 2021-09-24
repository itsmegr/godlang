package main

import "fmt"
/*
	normal flow of execution
	if panic is not there
		1. all statements
		2. all defered functions
		return
	if panic is there
		1. till the statements where panic occured
		2. all defered functions

		if panic is recovered return to the calling function and continue execution
		else terminate the programme
*/

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("medium")
	fmt.Println("end")
}

func panicDemo() {
	a, b := 10, 0;
	defer func(){
		err1 := recover();
		fmt.Println(err1);
	}()
	fmt.Println(a/b);
}

func main(){
	/*
    	defer keyword --> execute the statement before returning the function
		if more than one defer than order will (first in last out) stack way
	*/
	// deferDemo()

	/*
		panic ---> its exception in go---> stops execution of programe and go to deferred funcs
		panic === throw
		recover === catch
	*/
	panicDemo()
	fmt.Println("it is recoverd");	

}