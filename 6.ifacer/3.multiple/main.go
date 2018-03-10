package main

import "fmt"

// PatientDog knows how to behave
type PatientDog interface {
	Sit()
	Stay()
}

// HyperDog can only behave for a short time
type HyperDog interface {
	Sit()
	Speak()
}

// SuperCorgi is the ultimate in Corgi
type SuperCorgi struct {
	PatientDog
	HyperDog
}

// PembrokeBehaviors defines a Pembrokes
type PembrokeBehaviors struct{}

// Sit will sit
func (p *PembrokeBehaviors) Sit() {
	fmt.Println("Ok, Pembroke is Sitting!")
}

// Stay will stay
func (p *PembrokeBehaviors) Stay() {
	fmt.Println("Ok, Pembroke is Staying!")
}

// CardiganBehaviors defines a Cardigan
type CardiganBehaviors struct{}

// Sit will sit
func (g *CardiganBehaviors) Sit() {
	fmt.Println("I'm a Cardigan!!")
}

// Speak will speak
func (g *CardiganBehaviors) Speak() {
	fmt.Println("Seeya from a Cardigan!!")
}

// DoPatient runs through the Patient behavior list
func DoPatient(p PatientDog) {
	p.Sit()
	p.Stay()
}

// DoHyper runs through the Greeter behavior list
func DoHyper(h HyperDog) {
	h.Sit()
	h.Speak()
}

func main() {
	corg := &SuperCorgi{PatientDog: &PembrokeBehaviors{}, HyperDog: &CardiganBehaviors{}}
	//corg.Sit() // Compiler Error: ambiguous selector corg.Sit

	// But also! These next two won't work either!

	//DoPatient(corg) // This actually has 2 compiler errors, one of which cascades into the other!
	// 1: SuperCorgi.Sit is ambiguous. Due to promotion rules, Sit does not make it into the MethodSet
	// 2: cannot use *SuperCorgi as type PatientDog: *SuperCorgi does not implement PatientDog (missing Sit method)

	// DoHyper(corg) // This has the same two errors as before, but for HyperDog instead
	fmt.Println(corg)
}
