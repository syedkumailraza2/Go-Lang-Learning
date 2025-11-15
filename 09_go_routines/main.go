package main
import (
	"fmt"
	"time"
)

func task (id int){
	fmt.Println("Doing Task number:",id)
}

func main (){
	for i := 0; i <=10; i++ {
		go task (i)
	}

	time.Sleep(time.Second *2)
}