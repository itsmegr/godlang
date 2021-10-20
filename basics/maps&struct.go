package main

import (
	"fmt"
	"reflect"
)

/*
in go collection type includes :- array, slices, map, struct
to declare them common syntax
identifier := type{

}
var identifier type := type{

}
*/

//to export start key also with uppercase, with type name
type Student struct {
	name string
	rn int
	clg string
}


func main(){
	//maps :- key valur pair, just same like c++
	//if key is not there it will give 0 value
	//refrence is copied
	var m1 map[string]int = map[string]int{
		"rn" : 2,
		"rn2" : 5,
	}
	fmt.Println(m1["rn"], m1["rn3"])
	m1["rn3"] = 34;

	delete(m1, "rn");
	fmt.Println(m1["rn"], m1["rn3"])

	//check weather key is there or not, _ give the value, put _, because don't want to use
	_, ok := m1["rn"];
	fmt.Println(ok)

	fmt.Printf("%v\n", m1);


	//STRUCT

	std1 := Student{
		name : "govind",
		rn  : 12,
		clg : "nitw", 
	}


	fmt.Println(std1.name);
	//creating copy duplicates

	std2 := struct{name string; rn int}{
		name : "govinf2",
		rn :10,
	}
	fmt.Println(std2);

	//tag in struct --> providing extra info about key
	type Animal struct{
		name string `required max:"100"`
		origin string
	}

	//getting the type
	tp := reflect.TypeOf(Animal {})
	//getting the field
	field, _ := tp.FieldByName("name");
	fmt.Println(field.Tag)
}