package main

import (
	"fmt"
)

func main(){
	a := 10;
	guess := 10;
	//here this true false works same as c++
	if guess < a {
		fmt.Println("oohhh number is greater than this too")
	} else if guess > a {
		fmt.Println("oohhh number is smaller than this too")
	} else{
		fmt.Println("great you found the num ber")
	}

	//switch--> for more info refer gfg
	switch name:="b"; name {
	case "a":
		fmt.Println("its a")
		fmt.Println("its a")
		fmt.Println("its a")
	case "govind" :
		fmt.Println("great its govind")
		fmt.Println("great its govind")
		fmt.Println("great its govind")
	default :
		fmt.Println("its any other")
		fmt.Println("its any other")
		fmt.Println("its any other")
	}

	var i interface{} = []int{1, 3, 4};
	switch i.(type) {
	case int : 
		fmt.Println("this is int")
	case float64 : 
		fmt.Println("this is float64")
	case []int : 
		fmt.Println("this is slice")
	default : 
		fmt.Println("this is noone")
	}
	
	
	
	

}