package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan time.Duration)
	go bucle(chan1)
	fmt.Println("Llegue :3")

	msg := <-chan1
	fmt.Println(msg)
}

func bucle(chan1 chan time.Duration) {
	init := time.Now()
	for i := 0; i < 10000000000000; i++ {

	}
	final := time.Now()
	chan1 <- final.Sub(init)
}
