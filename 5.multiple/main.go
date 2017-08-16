package main

import (
	"fmt"
)

// Pembroke WelshCorgi has a name
type Pembroke struct {
	Name string
}

// Cardigan Corgi has a name
type Cardigan struct {
	Name string
}

// Hello says hello
func (a *Pembroke) Hello() {
	fmt.Println(a.Name)
}

// Hello also says hello
func (b *Cardigan) Hello() {
	fmt.Println(b.Name)
}

// SuperCorgi embeds PembrokeWelshCorgi and CardiganCorgi into one unstoppable corg
type SuperCorgi struct {
	Pembroke
	Cardigan
}

func main() {
	sc := SuperCorgi{Pembroke: Pembroke{Name: "Bark Ruffalo"}, Cardigan: Cardigan{Name: "Bark Hamill"}}

	// sc.Hello() 				// Compiler Error: ambiguous selector sc1.Hello
	// fmt.Println(sc.Name) 	// Compiler Error: ambiguous selector sc2.Name

	fmt.Println(sc)
}
