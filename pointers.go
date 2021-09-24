package main

import "fmt"

func main() {
	/*
		pointers = memory address of variables
		derefrecing operator = getting the value at that address
	*/

	var a int =14;
	var b *int = &a;
	fmt.Println(&a, b);
	fmt.Println(a, *b); //derefrencing

	//pointer arithmatic is not allowed in go
	//if no value is assigned to a pointer type its nil

	var a1 *st;
	a1 = new(st); //new keyword always gives the address
	// (*a1).name = 2
	// fmt.Println((*a1).name);
	//another way is just by this is understood by compiler
	a1.name = 2;
	fmt.Println(a1.name)


	//in array coompletely copy is created but in slices, map copy only refrence y

}

type st struct {
	name int
}