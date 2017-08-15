package main

import "fmt"

// GoodDogList is a type which is actually a []bool for tracking Good Dogs
type GoodDogList []bool

// CorgiSet is a struct which embeds GoodDogList
type CorgiSet struct {
	GoodDogList // Allowed!
}

// PointerCorgiSet is a struct which embeds *GoodDogList
type PointerCorgiSet struct {
	*GoodDogList // Also allowed!
}

func main() {
	gdl := make(GoodDogList, 1)
	fmt.Println(gdl[0]) // Works, because GoodDogList is a slice directly
	corgSet := CorgiSet{GoodDogList: gdl}
	// fmt.Println(corgSet[0]) // Compiler Error: invalid operation: corgiSet[0] (type CorgiSet does not support indexing)
	// This is because map and array/slice indexers do not get promoted!
	fmt.Println(corgSet)
}
