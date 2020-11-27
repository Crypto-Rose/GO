package main

import "fmt"

type livingBeing interface {
	itsAlive() bool
}

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
	itsCarnivore() bool
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
	alive     bool
}

type women struct {
	men
}

type dog struct {
	breathing bool
	eating    bool
	carnivore bool
	alive     bool
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
func (h *men) itsAlive() bool { return h.alive }

func (p *dog) breathe()           { p.breathing = true }
func (p *dog) eat()               { p.eating = true }
func (p *dog) itsCarnivore() bool { return p.carnivore }
func (p *dog) itsAlive() bool     { return p.alive }

func AnimalBreathing(an animal) {
	an.breathe()
	fmt.Printf("Soy un animal y estoy respirando")
}

func AnimalCarnivore(an animal) int {
	if an.itsCarnivore() == true {
		return 1
	}
	return 0
}

func HumanBreathing(hu human) {
	hu.breathe()
	fmt.Printf("Soy un/a %s, y estoy respirando \n", hu.gender())
}

func IAmAlive(v livingBeing) bool {
	return v.itsAlive()
}
func main() {
	totalCarnivore := 0
	Guffy := new(dog)
	Guffy.carnivore = true
	Guffy.alive = true
	AnimalBreathing(Guffy)

	totalCarnivore = +AnimalCarnivore(Guffy)
	fmt.Printf("Total carnivoros %d \n", totalCarnivore)
	fmt.Printf("Estoy vivo = %t", IAmAlive(Guffy))
	/*Daniel := new(men)
	Daniel.isMen = true
	HumanBreathing(Daniel)

	Rose := new(women)
	HumanBreathing(Rose)*/
}
