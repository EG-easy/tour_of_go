package basic2

import "fmt"

func main(){
	defer fmt.Println("world")
	fmt.Println("hello")
}
