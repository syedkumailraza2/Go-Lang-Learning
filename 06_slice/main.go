package main
import "fmt"

func main(){
	var mySlice = []int{1,2,3,4}
	fmt.Println(mySlice)

	//slice made from an array
	var arr = [...]int{12,32,42,54,65,54}
	var sliceFromArr = arr[0:3]
	fmt.Println(sliceFromArr)

	//creating slice using make() function
						//make([]type,length,capacity)
	var sliceUsingMake = make([]int, 3, 5)
	sliceUsingMake = append(sliceUsingMake, 21,32,54,65,43,32,32)
	fmt.Println(sliceUsingMake)
}
