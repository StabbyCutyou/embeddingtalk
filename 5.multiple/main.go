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
	sc1 := SuperCorgi{Pembroke: Pembroke{Name: "Pembroke1"}, Cardigan: Cardigan{Name: "Cardigan1"}}
	sc2 := SuperCorgi{Cardigan: Cardigan{Name: "Cardigan2"}, Pembroke: Pembroke{Name: "Pembroke2"}}

	// sc1.Hello() 				// Compiler Error: ambiguous selector sc1.Hello
	// fmt.Println(sc2.Name) 	// Compiler Error: ambiguous selector sc2.Name

	fmt.Printf("%+v\n", sc1)
	fmt.Printf("%+v\n", sc2)
}
