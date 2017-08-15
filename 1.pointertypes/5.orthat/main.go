package main

// GoodDogList is a type which is actually a []bool for tracking Good Dogs
type GoodDogList []bool

// PointerGoodDogList is a pointer to a type which is actually a []bool for tracking Good Dogs
type PointerGoodDogList *GoodDogList

// CorgiSet is a struct which embeds PointerGoodDogList
type CorgiSet struct {
	PointerGoodDogList // Compiler Error: embedded type cannot be a pointer
}

func main() {

}
