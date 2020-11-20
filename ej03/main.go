package main

import "fmt"

var state bool

func main() {
	state = true
	if state = false; state == true {
		fmt.Println(state)
	} else {
		fmt.Println("This is a false")
	}

	switch number := 5; number {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	default:
		fmt.Println("Mayor a 5")
	}
}
