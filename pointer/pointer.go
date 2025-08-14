package main

import "fmt"

func main () {
	a := 10;
	//address of "a" in ram
	b := &a // 0xc00000a0d8
	// access value of "0xc00000a0d8" which is address of "a" in ram
	c := *b //10
	fmt.Println(c)
}