package main

import "fmt"

/*
const keyword is used for it
--- typed const => where type is specified
--- untyped where no type is specified and compiler decides
--- const replaces its idetifier with the literal value
** value must be calculable at compile time

** name them same like variable
*/


/*
enumeration consts
each const is mapped with a value
*/

const (
	_= iota
	isAdmin = 1 << iota //iota is equal to zero but its special keyword and each value will be left shift of previuos
	isStudent
	isTeacher
)

func main(){
	const v1  = 12;
	fmt.Printf("%v, %T\n", v1, v1);

	const v2 int32  = 1000;
	fmt.Printf("%v, %T\n", v2, v2);


	const user1 = isAdmin | isTeacher;
	fmt.Printf("isAdmin? %v\n", user1 & isAdmin == isAdmin)
	fmt.Printf("isAdmin? %v\n", user1 & isStudent == isStudent)
}