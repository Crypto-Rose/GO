package main

import (
	"fmt"
	"strconv"
)

var number int
var text string
var status bool = true

func main() {
	number2, number4, text, status := 3, 4, "'My Text", false
	//Es una variable unica como name:=10

	var moneda int64 = 0
	number2 = int(moneda)
	//text = fmt.Sprintf("%d", moneda)
	text = strconv.Itoa(int(moneda))

	fmt.Println(number2)
	fmt.Println(number4)
	fmt.Println(text)
	fmt.Println(status)
	//Para exportar se utilizan las Capitalize o se coloca la letra primera en mayuscula
	MostrarStatus()

}

func MostrarStatus() {
	fmt.Println(status)
}
