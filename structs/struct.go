package main

import (
"time"
"fmt"	
)

// if we have to make complex data structure with multiple fields

// create a struct for e commerce application

// type order struct {         // if you dont set values , then the default value is set which is zero value 
// 	id string
// 	amount float32
// 	status string
// 	createdAt time.Time // its a nanosecond precision 
// }


// // how to attach methods in struct 

// func (o *order) changeStatus(status string){  // o order is receiver type , which help us to attach the method to order struct , but we also have to pass the pointer

// 	o.status=status; /// struct automatically dereference this so no need of * 

// }




// how to create a constructor for the struct 

type order struct {
	id string
	amount float32
	status string
	createdAt time.Time
	customer    //  this is how struct embedding is done 
}
type customer struct {
	name string
	phone string
}

// this is like a constructor 

func newOrder(id string, amount float32, status string) *order{  // return value is of order struct 
	myOrder:=order{
		id: id,
		amount:amount,
		status:status,
	}


	return & myOrder;   // we dont return struct directly , we return a pointer of it 
}



// NOW HOW TO REFERENCE ONE STRUCT INSIDE OTHER SUCH AS WE DO IN MONGO DB refreencing types object 
func main(){

	order:= order {
		id:"1",
		amount:450,
		status :"received",

	}


	// order.changeStatus("confirmed");

	fmt.Println(order);

}