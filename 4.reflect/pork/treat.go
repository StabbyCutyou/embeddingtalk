package pork

import "fmt"

// Treat is a tasty pork treat for dogs to eat
type Treat struct {
	size        int32
	SourcedFrom string
}

// DriedEar sounds gross, but dogs go nuts for'em
type DriedEar struct {
	size        int32
	SourcedFrom string
}

// Description will describe itself
func (t *Treat) Description() {
	fmt.Printf("Hi, i'm a size %d Pork Treat\n", t.size)
}

// Description will say describe itself
func (e *DriedEar) Description() {
	fmt.Printf("Hi, i'm a size %d Pork Ear\n", e.size)
}
