package main

import (
	"fmt"

	"github.com/StabbyCutyou/embeddingtalk/2.methodsets/pkg"
)

type E struct {
	name string
}

// SayHi will say hellos
func (e E) SayHi() {
	fmt.Println("Hi, i'm" + e.name)
}

func (e E) hello() {
	fmt.Println("Hey")
}

// S is a struct that embeds E, with unexported fields from the same package
type S struct {
	E
}

// G is a struct that embeds pkg.T, with unexported fields from another package
type G struct {
	pkg.T
}

func main() {
	s := S{E: E{name: "alex"}}
	//g := G{T: pkg.T{name: "davida"}} // Cannot access unexported fields across packages!
	g := G{T: pkg.T{}}
	s.SayHi()
	g.SayHi()

	s.hello()
	// g.hello() // unexported `hello` cannot be embedded across packages!
}
