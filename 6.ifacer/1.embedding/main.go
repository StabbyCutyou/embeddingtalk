package main

import (
	"fmt"
)

// Barker is an interface
type Barker interface {
	Bark() bool
}

// NoiseBox enables a Corgi to make noise
type NoiseBox struct{}

// Bark satisfies Barker
func (n *NoiseBox) Bark() bool {
	fmt.Println("Bark, yall!")
	return true
}

// Corgi is a struct
type Corgi struct {
	Barker
}

func main() {
	c := &Corgi{Barker: &NoiseBox{}}
	// Bark is called on the underlying NoiseBox by way of Barker
	fmt.Println(c.Bark())
	c = &Corgi{}
	// Bark will panic! Because Barker is nil!
	fmt.Println(c.Bark()) // panic!
}


