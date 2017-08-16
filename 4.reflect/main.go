package main

import (
	"fmt"
	"reflect"

	"github.com/StabbyCutyou/embeddingtalk/4.reflect/lamb"
	"github.com/StabbyCutyou/embeddingtalk/4.reflect/pork"
)

// Snack is a tasty store bought treat, built to spec
type Snack struct {
	size int
	pork.DriedEar
	lamb.Treat
}

// HomemadeSnack is made with love, but the final product never has a well defined size
type HomemadeSnack struct {
	lamb.Treat
	pork.DriedEar
}

func main() {
	s := Snack{}

	fmt.Println(reflect.TypeOf(s.size)) // int, not int32 nor int64
	_, ok := reflect.TypeOf(s).FieldByName("size")
	fmt.Println(ok) // Since the size field is native to the struct, this prints `true`

	// I can make my own snack, better than those other guys!
	hs := HomemadeSnack{} // Has no native size

	_, ok = reflect.TypeOf(hs).FieldByName("size")
	fmt.Println(ok) // It's not native to the struct, thus it can't be found, this prints `false`

	_, ok = reflect.TypeOf(hs).MethodByName("Description")
	fmt.Println(ok) // It's ambiguous, thus the method is not part of the MethodSet, this prints `false`
}
