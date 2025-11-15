package main

import "fmt"

func main(){
	var arr1 = [3]int{1,2,3} //normal declaration with var keyword
	arr2 := [4]int{1,2,3,4} //shortcut declaration using := sign
	arr3 := [...]int{2,4,5,6,45,6,4,5,3} //here the lenght is inferred (length will be assigned dynamically)

	//printing all the values of array
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	//accessing values in an array
	fmt.Println(arr3[2])
	fmt.Println(arr3[5])

	//changing element in an array
	fmt.Println("Before changing")
	fmt.Println(arr3)
	arr3[2] = 32
	fmt.Println("After changing")
	fmt.Println(arr3)

	//initializing an array
	var arrInit1 = [4]int{} //not initialized
	var arrInit2 = [12]int{2,4,3} // partially initialized
	var arrInit3 = [3]int{4,12,43} // fully initialized

	//finding length of an array
	fmt.Println(len(arrInit2))
}	