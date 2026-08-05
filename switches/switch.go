package main

import (
	"fmt"
	// "time"
)

func main(){

// i:=5;
// // this is the simple type of switch 
// switch i{

// case 1:
// 	fmt.Println("one");
// case 2:
// 	fmt.Println("two");
// case 3:
// 	fmt.Println("three");
// case 4:
// 	fmt.Println("four");
// case 5:
// 	fmt.Println("five");

// default:
// 	fmt.Println("others")    // automatic handles break behind the scene
// }



/// this is the multipe condition switch

// switch time.Now().Weekday(){

// case time.Sunday , time.Saturday:
// 	fmt.Println("this is weekend")

// default:
// 	fmt.Println("its work day");
// } 






// THIS IS THE TYPE SWITCH 

 whoAmI:= func(i interface{}){
	switch t:= i.(type){

	case int:
		fmt.Println("Its an integer");
	case string:
		fmt.Println("Its an string");
	case bool:
		fmt.Println("Its an boolean type");
	default:
		fmt.Println("other",t);
	}
 }


 whoAmI("golang");
 whoAmI(1);
}