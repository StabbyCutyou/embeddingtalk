package pork

import "fmt"

// Treat is a struct we export
type Treat struct {
	name        string
	SourcedFrom string
}

// SayHi will say hellos
func (t Treat) SayHi() {
	fmt.Println("Hi, i'm" + t.name)
}

func (t Treat) hello() {
	fmt.Println("Hey")
}

// DriedEar is a struct we export
type DriedEar struct {
	name        string
	SourcedFrom string
}

// SayHi will say hellos
func (e DriedEar) SayHi() {
	fmt.Println("Hi, i'm" + e.name)
}

func (e DriedEar) hello() {
	fmt.Println("Hey")
}
