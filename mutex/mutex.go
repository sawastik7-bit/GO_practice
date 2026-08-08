// package main

// import (
// 	"fmt"
// 	"sync"
// )

// // mutual exclusion ,, // to prevent from Race condition

// // what is race conditon , when multiple processes modifying same resource, then it wont be consistent or atomic

// type post struct{
// 	views int

// }

// func(p *post) Increment(wg *sync.WaitGroup){
// 	defer wg.Done();
// p.views+=1;

// }

// func main(){

// 	var wg sync.WaitGroup;

// mypost:=post{views: 0};

// for i:=0;i<100;i++{
// 	wg.Add(1);
// 	go mypost.Increment(&wg);  // it wont increment sequentially , it will run concurrently
// }

// wg.Wait();
// fmt.Println(mypost.views);

// }

// First we have to create a mutex

package main

import (
	"fmt"
	"sync"
)



type post struct{
	views int 
	mu sync.Mutex        // now in increment function before view modification do p.mu.lock() and after modificagtion , do p.mu.unlock()
}

func main(){
fmt.Print()
}