package lamb

import "fmt"

// Treat is a tasty lamb treat for dogs to enjoy
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
