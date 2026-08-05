package main

// import "fmt";

// func add(a , b int) int{

// 	return a+b;
// }


// func getLanguages() (string,string,string){  /// for multiple return types
// 	return "golang","js","cpp"
// }



func Process(fn func(a int) int){
	fn(1);
}

func main(){ // this is also a function , built with func

// // var ans=add(3,4);

// // fmt.Println(ans);

// lang1, lang2,lang3:=getLanguages();
// fmt.Println(lang1,lang2,lang3);

// // we can also do 
// lang1,lang2,_=getLanguages(); // the _ will tell us that we are returning some value but not using it 

fn:=func(a int)int{
	return 2;
}
Process(fn);
}