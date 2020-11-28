package main

import "fmt"

//Recursion
func main() {
	exponential(2)
}

func exponential(number int) {
	if number > 100000000 {
		return
	}
	fmt.Println(number)
	exponential(number * 5)
}
