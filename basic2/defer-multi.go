package basic2

import "fmt"

func main(){
	fmt.Println("counting")

	for i := 0; i < 10; i++{
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

