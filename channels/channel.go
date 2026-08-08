// package main

// import (
// 	"crypto/rand"
// 	"fmt"
// 	"time"
// )

// // channels are like pipe , where we send the data and receive the data
// // if we run multipe go routines which are concurrent , and if we want to send one ddata from one go routine to other then we send them through channels

// func processNum(numChan chan int){

// 	for num:= range numChan{
// 		fmt.Println("Processing number ", num);

// 		time.Sleep(time.Second*1);
// 	}
// }


// func main(){


// 	// numChan:= make(chan int);

// 	// go processNum(numChan);

// 	// numChan <-5;

// 	// time.Sleep(time.Second *2);

// 	// Now for an infinite go routine channel pipeline 

// 	numChan:=make (chan int );

// 	go processNum(numChan);

// 	for{
// 		numChan <- rand.Intn(100);
// 	}

	

// // 	messageChannel:=make(chan string) // the type of data you want to send 

// // 	messageChannel <- "ping"  // the direction to which you want to send the data 

// // msg:= <- messageChannel

// // 	fmt.Println(msg);

// }

package game