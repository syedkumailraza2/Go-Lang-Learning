package main

import "fmt"

func main(){
	var studentName string = "Raza"
	var studentAge = 21 //this is a inferred type of declaration the type of the variable declared based on the value

	//this is the declaration of variables without assigning them
	var a string
	var x int
	var c bool

	//assigning values after declaration
	a = "Hello"
	x = 23
	c = true

	fmt.Println("StudentName: " + studentName)
	fmt.Println(studentAge)
}