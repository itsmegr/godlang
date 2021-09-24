package main

import (
	"fmt"
	"strconv"
)

/*
param1 = same like variable declaration but without var keyword
func name(param1, param2)returnTyope{
	body
}
1. paramthesis shoould start with the same line
*/

func printName(fn string, ln string, rn int){
	fmt.Println(fn, ln, rn)
}

func printName2(fn, ln string, rn int){
	fmt.Println(fn, ln, rn)
}

func printName3(rn int, fullName ...string) string{
	retValue := fullName[0] + " " + fullName[1] + " " +  strconv.Itoa(rn);
	return retValue;
}

//naming return value
func printName4(rn int, fullName ...string) (fullData string){
	fullData = fullName[0] + " " + fullName[1] + " " +  strconv.Itoa(rn);
	return 
}


func devide(a ,b float64) (float64, error) {
	if b==0 {
		return 0.0, fmt.Errorf("B cannot be zero")
	}

	return a/b, nil
}


func main() {
	printName("govind", "rathore", 30)
	printName2("govind", "rathore", 30)
	//using return
	fmt.Println(printName3(30, "govind", "rathore"))

	//using naming return
	fmt.Println(printName4(30, "govind", "rathore"))

	//in golang a function can return multiple arguments
	res, err1 := devide(12,0);
	if err1 !=nil{
		fmt.Printf("%v is type of %T\n", err1,err1);
	} else{
		fmt.Println(res)
	}

	//we can treat functions like variable and can do all the stuff that we cann do with functions
	var func1 func() = func() {
		fmt.Println("hello from anonymous function")
	}

	func1();

	//functions as method ---> just like methods asscociated with any type
	/*
		greeter is a struct
		func (localName TypeOfObject) funcName(params){
			//here goes the body
		}
		
	*/
}
