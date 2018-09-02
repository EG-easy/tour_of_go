package basic2

import (
	"time"
	"fmt"
)

func main(){
	t := time.Now()
	switch{
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() <17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening.")
	}
}
