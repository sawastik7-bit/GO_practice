package main

import "fmt"



func sum(nums ...int) int{  // for any type we can use interface{} instead of ...int

	total:=0;

	for _,num:=range nums{
		total=total+num
	}

	return total;

}

func main(){

	result:=sum(1,2,3,4,5,6);  //a function is called variadic if it receives n number of parameters

	nums:=[]int{3,4,5,6};
	result2:=sum(nums...);  // to pass a slice instead of seperate values 
	fmt.Println(result);

	fmt.Println(result2);


}