package main

import "fmt"

type Greeter interface {
	Hello()
	Goodbye()
}

type CasualGreeter interface {
	Hello()
	Seeya()
}

type Goodbyeer interface {
	Goodbye()
	Seeya()
}

type GoodDog struct {
	Greeter
	CasualGreeter
}

type PembrokeBehaviors struct{}

func (p *PembrokeBehaviors) Hello() {
	fmt.Println("I'm H!")
}
func (p *PembrokeBehaviors) Goodbye() {
	fmt.Println("Goodbye from H!")
}

type CardiganBehaviors struct{}

func (g *CardiganBehaviors) Hello() {
	fmt.Println("I'm G!")
}
func (g *CardiganBehaviors) Seeya() {
	fmt.Println("Seeya from G!")
}

func DoGreeter(g Greeter) {
	g.Hello()
	g.Goodbye()
}

func DoCasualGreeter(cg CasualGreeter) {
	cg.Hello()
	cg.Seeya()
}

func DoGoodbyeer(gb Goodbyeer) {
	gb.Goodbye()
	gb.Seeya()
}

func main() {
	gd := &GoodDog{Greeter: &PembrokeBehaviors{}, CasualGreeter: &CardiganBehaviors{}}
	// gd.Hello() // Compiler Error: ambiguous selector gd.Hello

	// But also! These next two won't work either!

	// DoGreeter(gd) // This actually has 2 compiler errors!
	// 1: GoodDog.Hello is ambiguous. This is due to promotion, which means Hello does not make it into the MethodSet!
	// 2: cannot use *GoodDog as type Greeter in argument: *GoodDog doesn't implement Greeter (missing Hello method)

	// This has the same two errors as before, but for CasualGreeter instead
	// DoCasualGreeter(gd)
	DoGoodbyeer(gd) // This is fine, because it avoids the ambiguous selector!
}
