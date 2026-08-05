package main

import "fmt"

// maps -> hash, object, dictionary


func main(){


	// creating map

	m:=make(map[string]string);

	// adding elements 
	m["one"]="rajesh";

	fmt.Println(m); // to print whole key value pairs
	fmt.Println(m["one"]); //// to print the specific value 

	fmt.Println(m["gaming"]);  /// if we try to fetch a non existing key in map , then it will return a 0 value , or string according to value type defined in map

	delete(m,"one"); // delete a specific key from map 



	// How to make map without make 

	newmap:= map[string]int{"price":34000,"phone":43};  // if we know elements from starting 

	fmt.Println(newmap);


	// how to check if element is there in map or not or to have to take actions on it 

	_,ok:=newmap["price"];  // its like a structure to check if something is present in  map or not 

	if(ok){
		fmt.Println("ALL OK");
	}else{
		fmt.Println("not okay");


	}

	v,ok:=newmap["price"]; // v represents the value that the current key is holding 
	fmt.Println(v);

}