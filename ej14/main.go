package main

//"io/ioutl"
//"log"

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
	a := 1
	if a == 1 {
		panic("Find the valor 1")
	}
}
