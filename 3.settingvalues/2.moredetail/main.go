package main

import (
	"fmt"
)

// G is a struct
type G struct {
	Name string
}

// SayMyName says my name
func (g *G) SayMyName() {
	fmt.Println(g.Name)
}

// S embeds G
type S struct {
	G
	Name string
}

// SayMyName2 says my name, too
func (s *S) SayMyName2() {
	fmt.Println(s.Name)
}

// P embeds *G
type P struct {
	*G
	Name string
}

// SayMyName says my name
func (p *P) SayMyName() {
	fmt.Println(p.Name)
}

func main() {
	s := S{
		Name: "Joe",
		G:    G{},
	}
	s2 := S{
		Name: "Joe",
		G: G{
			Name: "Sam",
		},
	}
	s3 := s2
	s3.Name = "Lisa"

	s4 := s2
	s4.G.Name = "Nicky"

	p := P{
		Name: "Joe",
		G:    &G{},
	}
	p.G.Name = "Nicky"

	fmt.Printf("s  : %+v\n", s)
	s.SayMyName()  // This one is promoted from G
	s.SayMyName2() // This one is native to S

	fmt.Printf("s2 : %+v\n", s2)
	s2.SayMyName()  // This one is promoted from G
	s2.SayMyName2() // This one is native to S

	fmt.Printf("s3 : %+v\n", s3)
	s3.SayMyName()  // This one is promoted from G
	s3.SayMyName2() // This one is native to S

	fmt.Printf("s4 : %+v\n", s4)
	s4.SayMyName()  // This one is promoted from G
	s4.SayMyName2() // This one is native to S
	fmt.Printf("p  : %v\n", p)
	fmt.Printf("p.G: %+v\n", *p.G)
	p.SayMyName()
}
