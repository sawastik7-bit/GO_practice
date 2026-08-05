package main

import "fmt"

func main() {
	// for is the only one loop in go with different syntax

	// go dont have while keyword

	i := 1
	// THIS IS THE WHILE LOOP
	for i <= 3 {
		fmt.Println(i);
		i=i+1;
	}



	// if we have to do infinite loop 

	// for {

	// }


	// now the classic for loop 

	for i:=0;i<3;i++ {

		if(i==2){ break;  // this is the break statement 
		}
		fmt.Println("hello ");

	}



	// if i have to do something for some number of times then we can use range also 

	for i:= range 3{
		fmt.Println(i);  // will print upto 2
	}

}