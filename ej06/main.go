package main

import "fmt"

func main() {
	/*fmt.Println(one(8))
	number, state := two(2)
	fmt.Println(number)
	fmt.Println(state)*/
	fmt.Println(Calcule(10, 50))
	fmt.Println(Calcule(20, 60, 1, 7))
	fmt.Println(Calcule(30))
	fmt.Println(Calcule(40, 80, 6, 898, 1213, 5657, 3345, 1))

}

func one(number int) (z int) {
	z = number * 4
	return
}

func two(number int) (int, bool) {
	if number == 1 {
		return 5, false
	} else {
		return 8, true
	}
}

func Calcule(number ...int) int {
	total := 0
	//for _, num := range number {
	for item, num := range number {
		total = total + num
		fmt.Printf("\n Item %d \n", item)
	}
	return total
}
