package main

import "fmt"

func main() {
	//Un mapa es una estructura que se puede serializar, comoo un vector. El indice puede ser clave y valor
	/*countries := make(map[string]string, 5)
	fmt.Println(countries)

	countries["Colombia"] = "Bogota"
	countries["Mexico"] = "D.F"*/

	championship := map[string]int{
		"Colombia": 80,
		"Peru":     75,
		"Mexico":   68,
		"Brasil":   91}

	championship["Argentina"] = 85
	championship["Brasil"] = 70
	delete(championship, "Brasil")
	fmt.Println(championship)

	for team, point := range championship {
		fmt.Printf("El equippo %s, tiene un puntaje de: %d \n", team, point)
	}
	point, exist := championship["Argentina"]
	fmt.Printf("El puntaje capturado es: %d, y el equipo existe %t", point, exist)
}
