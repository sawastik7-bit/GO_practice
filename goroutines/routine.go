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

// waitgroups is done to synchronize go routines

package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup ){  // WAITGROUP SHOULD BE POINTER
	
	defer w.Done();
	
	fmt.Println("The task is running on id :", id);

}

func main(){

	// how to add waitgroups 
	var wg sync.WaitGroup

	for i:=0;i<10;i++{
		wg.Add(1);
go task(i,&wg);
	}

	wg.Wait();
}
