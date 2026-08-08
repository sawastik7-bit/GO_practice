package main

import (
	"fmt"
	"os"
)



func main(){

	// how to read a file 

	// the most used method is 

	// F,err:=os.Open("example.txt");


	// if err !=nil {
	// 	// log the error 
	// 	// or panic the program

	// 	panic(err);
	// }

	// fileinfo,err:=F.Stat();

	// if err !=nil {
	// 	// log the error 
	// 	// or panic the program

	// 	panic(err);
	// }

	// fmt.Println(fileinfo.Size());  // will show in bytes

	



	/// Now we will read through file 

	// f, err:=os.Open("example.txt");

	// if err !=nil{
	// 	panic(err);
	// }


	// defer f.Close();  // will run at the end 


	// // how to read a file 
	// buf :=make([]byte, 20);

	// d,err:= f.Read(buf);

	// if err!=nil{
	// 	panic(err);
	// }

	// for i:=0 ;i<len(buf);i++{
	// 	println("data in the buffer ",d,string(buf[i]));
	// }


	// Now how to read a file 

	// f, err:=os.ReadFile("example.txt");  // its not a viable solution cause it loads the entire module at once 

	// if err!=nil{
	// 	panic(err);
	// }

	// fmt.Println(string(f));




	// How to read folders 

	dir,err:=os.Open("../");

	if err!=nil{
		panic(err);
	}

	defer dir.Close();

	fileinfo, err:= dir.ReadDir(-1);  // will read the directories  // giving -ve will give the list of all of the directories 

	for _,fi:=range fileinfo{
		fmt.Println(fi.Name())
	}
	
}