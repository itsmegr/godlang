package main

import (
	"fmt"
)

/*
interfaces in Go --> concept wise they are same as java
implemetation is different
*/

//defining interfaces --> will have only functions and these functions are later
//can be associated with wny type

type classOps interface {
	pass() int
	fail() int
}

//lets say i have a struct
type student struct {
	curClass int
}

//now i want to associate that pass/fail function with this type
func (std *student) pass() int {
	std.curClass++;
	return std.curClass
}
func (std student) fail() int {
	std.curClass--;
	return std.curClass
}


//so till now that increment function is got associated with class type

func main(){
	//here the main thing comes, we implemented pass/fail with type student also
	std1 := student{5};
	std1.pass()
	//now we can assign student type to std1
	var op1 classOps = &std1
	fmt.Println(op1.pass())
	fmt.Println(std1.curClass)

	//we can also compose more than 1 interfaces and make 1 from them

	//type conversion

	//converting from interface to its type
	var std2 *student = op1.(*student)
	//now this std2 has access to all members and methods too
	fmt.Println(std2.fail())
	fmt.Println(std2.curClass)

	//methods set for a value type is all the methods having value as a receiver
	//methods set for a pointer type is all the methods having value & pointer as a receiver(means all)

}