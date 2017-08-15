package main

// Treat is a delicious thing for dogs to eat
type Treat struct{}

// TreatChannel is a type which is a channel of Treats for the dog to eat
type TreatChannel chan Treat

// Corgi is a struct that has a TreatChannel for when it is behaving like a good dog
type Corgi struct {
	TreatChannel
}

func main() {
	tc := make(TreatChannel, 1)
	c := Corgi{TreatChannel: tc}
	<-c  // Compiler Error: invalid opertion: <-c (receive from non-chan type Corgi)
	<-tc // no error! Channel operators do not get promoted, much like array and map indexers
}
