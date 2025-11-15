package main
import "fmt"

func sum(nums ...int) int {
	temp := 0
	for _, num := range nums {
		temp = temp + num
	}
	return temp
}

func main (){
	result := sum(1,32,432,54,6436,63346,657674)
	fmt.Println(result)
}