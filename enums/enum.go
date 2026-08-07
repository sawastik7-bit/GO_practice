package main

import "fmt"
                // we dont have enums in go 

				// we implement enum using consts 

				const (
					image int=45
					ranges int =66
				)


			type OrderStatus int  // we can create our own types 

		const (
			Received OrderStatus=45

		)

func changeOrderStatus(status OrderStatus){
	fmt.Println("changing order status to : ",status);   // but what if we have multiple status 

}

func main(){
 
	changeOrderStatus(Received);  // and we are here passing the string and may give typo error if there is some alphabetical mistake 

}

