package main

import (
	//Pintar en consola
	"fmt"
	//Para el sistema operativo
	"os"
	//Alternativa
	"bufio"
)

var number1 int
var number2 int
var result int
var leyend string

func main() {
	fmt.Println("Ingrese primer numero: ")
	fmt.Scanln(&number1)

	fmt.Println("Ingrese segundo numero: ")
	fmt.Scanln(&number2)

	fmt.Println("Description: ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		leyend = scanner.Text()
		//fmt.Printf("%d", number1+number2)
	}

	result = number1 + number2
	fmt.Println(leyend, result)

}
