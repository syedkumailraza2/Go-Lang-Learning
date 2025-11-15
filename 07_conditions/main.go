package main

import "fmt"

func main(){
	var num int = 23
	
	if num <= 18{
		fmt.Println("You are not eligible to vote")
	} else{
		fmt.Println("You are eligible to vote")
	}
}