package main

import (
	"fmt"
	"slices"
)
func main(){

	// slices automatically expand and contract 
	// These are dynamic arrays

	// mostly used construct in go 
	// useful methods in slices


	// we dont give the slice 
	// var nums []int;  // uninitialized slice is nill , which is in case of this 

	// fmt.Println(nums);


	// if we dont want to make the slice nill by default

// 	var nums=make([]int,2,5);   // for this slice the initial size is 2 so it would have [0,0],, the third parameter is initial capacity
// nums=append(nums,2);
// var nums2=make([]int,len(nums));



// copy(nums2,nums);  // destination and source as params

// fmt.Println(nums,nums2);
	
	// nums=append(nums,1);
	// 	fmt.Println(nums);
	// fmt.Println(cap(nums));// cap is capacity


	// // one another way to make a slice 

	// num:=[]int{};

	// fmt.Println(cap(num)); 


	// SLICE OPERATOR

	var nums=[]int {1,2,3};

	fmt.Println(nums[0:2]);


	// THERe is also a slice package

	var game=[]int {1,2};
	var game2=[]int {1,2};

	fmt.Println(slices.Equal(game,game2));  // compare in increasing fashion



	// We can also make 2d slices

	var game3=[][]int{{1,2,3},{4,5,6}};
fmt.Println(game3);

}

