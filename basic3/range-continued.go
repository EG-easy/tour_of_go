package main

import "fmt"

func main(){
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // ==2**1 <<はビット演算子を表す 101 << 1 → 1010(二進数)で十進数だと10
	}
	for _, value := range pow{
		fmt.Printf("%d\n", value)
	}
}


