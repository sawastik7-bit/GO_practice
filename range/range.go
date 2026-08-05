package main

import "fmt"

func main() {

	nums := []int{6, 7, 8}

	for _, num := range nums { //num represent the current number
		fmt.Println(num);
	}

	// now iterating over maps

	m:=map[string]string{"fname":"john","lname":"doe"};

	for k,v:=range m{
		fmt.Println(k,v);   // if we only want keys then only give one parameter 

	}


	// Range can also be used in strings

	for i,c:= range "going"{
		fmt.Println(i,c);   // it will print the index and unicode for the character  , its like code point rune 

		// here i is the starting byte of the rune 
	}
}


