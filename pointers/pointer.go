package main

import "fmt";


// func changeNum(num int){  // we have to pass this number by reference 

// 	num=5;

// 	fmt.Println("in changeNum ",num);
// }


func ChangeNum(num *int){
	*num=5;

}

func main(){
num:=1;


fmt.Println("the address of num is ", &num);

// ChangeNum(num);  // here the number was called by value so there was a copy in the function 
// ChangeNum(&num); // here its passed by reference

// fmt.Println("after change num ",num);  // here why the number didnt changes
 


}