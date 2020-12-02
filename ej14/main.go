package main

import (
	"log"
)

func main() {
	//Differ: Es una insctruccion que se ejecuta cuando hay un error o un return...
	/*file := "test.txt"
	f, err := os.Open(file)

	defer f.Close()

	if err != nil {
		fmt.Println("Error abrierdo el archivo")
		os.Exit(1)
	}*/

	examplePanic()
}

func examplePanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error que generó Panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Find the value 1")
	}
}

/*"fmt"
"io/ioutil"
"log"
"os"*/
