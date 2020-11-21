package main

import "fmt"

func main() {
	/*i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	number := 0
	for {
		fmt.Println("Continue")
		fmt.Println("Write secret number")
		fmt.Scanln(&number)
		if number == 100 {
			break
		}
	}

	var i = 0
	for i < 10 {
		fmt.Printf("\n Value of i: %d", i)
		if i == 5 {
			fmt.Printf(" multiply for 2 \n ")
			i = i * 2
			continue
		}
		fmt.Printf(" Paso por aqui \n")
		i++
	}*/

	var i int = 0
RUTINE:
	for i < 10 {
		if i == 2 || i == 4 {
			i = i * 2
			fmt.Println("Go to rutine,adding up 2!")
			goto RUTINE
		}
		fmt.Printf("Value of i: %d\n", i)
		i++
	}
}
