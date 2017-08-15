package main

import (
	"fmt"
	"reflect"

	"github.com/StabbyCutyou/embeddingtalk/4.reflect/lamb"
	"github.com/StabbyCutyou/embeddingtalk/4.reflect/pork"
)

// Snack is a struct which embeds things for dogs to eat
type Snack struct {
	// For some silly reason, this Snack uses integers for names!
	name int
	// pork.Treat // Compiler Error: duplicate field Treat - Leaving this uncommented marks lamb.Treat as duplicate!
	pork.DriedEar
	lamb.Treat
}

// HomemadeSnack is a struct which also embeds things for a dog to eat, but lacks a name
type HomemadeSnack struct {
	lamb.Treat
	// pork.Treat // Compiler Error: duplicate field Treat - Leaving this uncommented marks pork.Treat as duplicate!
	pork.DriedEar
}

func main() {
	s := Snack{name: 1}                 // Because of promotion semantics, the name field local to Snack wins over the others
	fmt.Println(reflect.TypeOf(s.name)) // int
	_, ok := reflect.TypeOf(s).FieldByName("name")
	fmt.Println(ok) // Since the name field is native to the struct, this prints `true`

	hs := HomemadeSnack{}
	// fmt.Println(reflect.TypeOf(g.name)) // Can't! name won't be exported!
	_, ok = reflect.TypeOf(hs).FieldByName("name")
	fmt.Println(ok) // It's not native to the struct, thus it can't be found, this prints `false`

	_, ok = reflect.TypeOf(hs).MethodByName("SayHi")
	fmt.Println(ok) // It's ambiguous, thus the method is not part of the MethodSet, this prints `false`

}
