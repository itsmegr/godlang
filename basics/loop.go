package main

import (
	"fmt"
)
func main(){
	for i:=1;i<=6;i++ {
		fmt.Println("heelo")
	}

	//in go 2 diffrent statements cannot be seperated by comma ","
	for i, j := 1,2 ; i < 6 && j < 5 ; i, j=i+2, j+2{
		fmt.Println(i, j);
	}

	i :=2;
	//same works like while
	for i <5 {
		if i > 3 {
			break;
		}
		fmt.Println(i);
		i++
	}
	i=2;
	//infinite loop 
	/*
	for {
		some statements
	}
	*/

	//concept of breaking to the the specific point, writing labels
	breakPoint :
	for {
		for i <5 {
			if i > 3 {
				break breakPoint;
			}
			fmt.Println("NOt broke", i);
			i++
		}
	}


	//using looop in collections --> using range

	coll := []int{1, 4, 5};
	for key, value := range coll {
		fmt.Println(key, value);
	}

	coll2 := map[int]string{
		1 : "govind",
		2 : "rathore",
	}
	for key, value := range coll2 {
		fmt.Println(key, value);
	}

	name := "govind";
	for _, value := range name {
		fmt.Println(string(value));
	}




}