package main

import "fmt"

//funciones anonimas- Se puedem modificar en tiempo de ejecucion
var Calcule func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("Sumo 5 + 7 =  %d \n", Calcule(5, 7))

	//Restamos
	Calcule = func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Printf("Resto 2-3 = %d \n", Calcule(2, 3))

	//Dividimos
	Calcule = func(num1 int, num2 int) int {
		return num1 / num2
	}
	fmt.Printf("Divido 9-3 = %d \n", Calcule(9, 3))

	Operations()

	/*	CLOSURES */

	tableDel := 2
	MyTable := Table(tableDel)
	for i := 0; i < 11; i++ {
		fmt.Println(MyTable())
	}
}

func Operations() {
	result := func() int {
		var a int = 30
		var b int = 10
		return a + b
	}
	fmt.Println(result())
}

//Closures: Concepto de funciones anonimas, tienen que ver con la proteccion de codigo

func Table(value int) func() int {
	number := value
	secuencie := 0
	return func() int {
		secuencie += 1
		return number * secuencie
	}
}
