package main

import "fmt"

/*
copy array, creates literal array to change this assign &(address)
*/

/*
slices is same as array but of dynamic size and can be sliced any part of that to another slice
*/

func main(){
	//ARRAY
	var grades [3]int = [...]int{1, 2, 3};//way of assigning value to array type
	// grades[0] = 1;
	// grades[1] = 2;

	grades2 := [...]int{1, 2, 3, 5}//these three dots makes its size to whatever is given after that, but do not make it dynamic

	fmt.Printf("%v, %T \n", grades, grades);
	fmt.Printf("%v, %T \n", grades2, grades2);
	fmt.Println(len(grades2))

	array1 := [...]string{"a", "b", "c"};
	array2 := &array1;
	array2[1] = "c";
	fmt.Printf("%v\n", array1);
	fmt.Printf("%v\n", array2);

	//SLICE
	a := []int{1, 2, 4, 5}
	fmt.Printf("%v\n", len(a));
	fmt.Printf("%v\n", cap(a))
	//this append function make the increases the capacity of slice
	a = append(a, 12, 21, 32, 43, 54, 65);
	// a = append(a, []int{12, 21, 32, 43, 54, 65, 44}...);//another way of appending
	fmt.Printf("%v\n", len(a));
	fmt.Printf("%v\n", cap(a))
	//slice is a refrence type so already refrence is assigned
	b := a;
	b[2]=10;
	fmt.Printf("%v\n", a);
	fmt.Printf("%v\n", b)


	//slicing the slice 🥱
	c := a[:]
	fmt.Printf("%v\n", c);
	d := a[3:]
	fmt.Printf("%v\n", d);
	//here index 2 is excluded
	f := a[:2];
	fmt.Printf("%v\n", f);

	e := make([]int, 4, 100);//type, size, capacity
	fmt.Printf("%v\n", e);
		
}