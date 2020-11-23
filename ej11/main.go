package main

import "fmt"

type human interface {
	//Se definen los metodos para implementar la interfaz
	breathe()
	think()
	eat()
	gender() string
}

type animal interface {
	breathe()
	eat()
	ItsCarnivore() bool
}

type vegetal interface {
	plantClassification() string
}

/*Human*/
type men struct {
	age       int
	height    float32
	weight    float32
	breathing bool
	thinking  bool
	eating    bool
	isMen     bool
}

type women struct {
	men
}

func (h *men) breathe() { h.breathing = true }
func (h *men) eat()     { h.eating = true }
func (h *men) think()   { h.thinking = true }
func (h *men) gender() string {
	if h.isMen == true {
		return "Men"
	} else {
		return "Women"
	}
}

func HumanBreathing(hu human) {
	hu.breathe()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", hu.gender())
}
func main() {
	Daniel := new(men)
	Daniel.isMen = true
	HumanBreathing(Daniel)

	Rose := new(women)
	HumanBreathing(Rose)
}
