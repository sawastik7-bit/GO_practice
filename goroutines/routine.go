// package main

// import (
// 	"fmt"
// 	"time"
// )

// // how to run concurrent processes

// func task(id int){
// 	fmt.Println("The task is running :",id);
// }

// func main(){

// for i:=0;i<=10;i++{
// 	go task(i);  // just write go in front of it , there are go routine lightweight threads , which will run concurrently , total 11 go routines parallely
// 	    // this is a blocking function ,, first one is printin then second then third ,,but how to run it parallely 
// }
//  // the main func should sleep 

//  time.Sleep(time.Second*2);  // will make the main func sleep till the go routine completes concurrency processes 
// }


// HERE WE WILL LEARN ABOUT waitgroups in go routine 

package main

import "fmt"

func main(){
fmt.Println();
}
