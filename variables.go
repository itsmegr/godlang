package main

import (
	"fmt"
	"strconv"
)

/*
on package level, variables can not be declared by : syntax
only way is default
*/
var fileName string = "variables";



/*
visibility of variables
1. if it start with lowercase, visible to package only(above fileName is this type)
2. if starts with Uppercase , than globally visible and file automatically exports it
3. block level variables (visible to blocklevel only eg. below name)
*/
var WhichLang string  = "golang"


func main(){
	//variable must be used after declaration 
	var i int = 10;
	fmt.Println(i);

	//type is detacted by compiler
	j := 13;
	fmt.Println(j);

	//here this need to be initialized on declaration only
	name := "govind rathore";
	fmt.Println(name);


	//type conversion
	//from int to float, float to int specify the function
	var rn int  = 14;
	var rn2 float32;

	rn2 = float32(rn);
	fmt.Printf("%T\n", rn2);

	//from int to string
	var rn3 string  = strconv.Itoa(14);
	fmt.Printf(rn3);

	//variable name length should be propotional to its life span, good practice

}