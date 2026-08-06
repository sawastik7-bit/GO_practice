package main
import "fmt"


// Interfaces are just contracts 

type paymenter  interface{    // add er in this 
 // we have to give contract here 

 pay(amount int);
}            




// payment integrate 
type Payment struct{

}


func (p Payment) makePayment(amount int){

	razorpayPaymentGateway :=razorpay{};

	razorpayPaymentGateway.pay(amount);

}


type razorpay struct{

}

func (r razorpay) pay(amount int){
	// logic to make payment

	fmt.Println("Making payment using razorpay instance ",25000)
}

func main(){

	Payment1 := Payment{};
	Payment1.makePayment(100);
}