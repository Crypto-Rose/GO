package main

import "fmt"

//Un slice es un vector dinamico el cual puedo ampliar las dimensiones en tiempo de ejecucion
//se crea como una matriz pero sin demensiones
//var table [15]int{8, 9, 50, 5, 2, 4, 1, 2, 5, 6}

func main() {
	//matrix := []int{2, 7, 6}
	//fmt.Println(matrix)
	variant2()
	variant3()
	variant4()
	/*table := [10]int{8, 9, 50, 5, 2, 4, 1, 2, 5, 6}
	for i := 0; i < len(table); i++ {
		fmt.Println(table[i])
	}*/
}

func variant2() {
	//3: Mostrar los datos desde la posicion 3 hasta la ultima
	elements := [5]int{2, 4, 6, 8, 5}
	portion := elements[3:]
	fmt.Println(portion)
}

func variant3() {
	//el make tiene 3 parametros, 1. El tipo de dato, 2. El tamaño 3. Capacidad
	elements := make([]int, 6, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(elements), cap(elements))
}

func variant4() {
	//Se llena un slice con append, este recibe un slice y retorna otro
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\n Lago %d, Capacidad %d", len(nums), cap(nums))
}
