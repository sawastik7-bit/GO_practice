package main

import "fmt"

func PrintSlice[T int|string](items []T) {
	for _, item := range items {
		fmt.Println(item);
	}
}

func main() {


	PrintSlice([]int{1,2,3,4,5});
	PrintSlice([]string{"rajesh","barish","burger"});


}