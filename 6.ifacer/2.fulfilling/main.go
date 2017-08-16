package main

import (
	"fmt"
)

// Barker is an interface for dogs
type Barker interface {
	Bark()
	Yelp()
}

// NoiseBox can bark
type NoiseBox struct{}

// Bark will bark
func (b *NoiseBox) Bark() {
	fmt.Println("Bark!")
}

// Yelp will yelp, but not on a pointer - thus, the interface is not fulfilled!
func (b NoiseBox) Yelp() {
	fmt.Println("Yipe yipe yipe!")
}

// Corgi is a special dog
type Corgi struct {
	NoiseBox
}

// DogStuff takes a Barker and does things
func DogStuff(b Barker) {
	b.Bark()
	b.Yelp()
}

func main() {
	c := Corgi{}
	// Does not work, because either way a method defines it's receiver
	// the embedded type still needs to be a pointer to fulfill the interface
	// DogStuff(c)
	DogStuff(&c)
}
