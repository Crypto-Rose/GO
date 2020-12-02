package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go myNameLasttt("Rosa morales")
	fmt.Println("estoy aqui")
	var x string
	fmt.Scanln(&x)
}

func myNameLasttt(name string) {
	letters := strings.Split(name, "")
	for _, letter := range letters {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letter)
	}
}
