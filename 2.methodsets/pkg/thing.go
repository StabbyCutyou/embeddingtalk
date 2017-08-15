package pkg

import "fmt"

// T is a struct we export
type T struct {
	name string
}

// SayHi will say hellos
func (t T) SayHi() {
	fmt.Println("Hi, i'm" + t.name)
}

// hello is not exported
func (t T) hello() {
	fmt.Println("Hey")
}
