package main

// CorgiSet is a struct that prevents compilation
type CorgiSet struct {
	// Keeps track of which Corgis are current Good Boys/Girls today
	[]bool // Compiler Error: syntax error: unexpected [, expecting field name or embedded type
}

// PointerCorgiSet is a struct that also prevents compilation
type PointerCorgiSet struct {
	// Keeps track of which Corgis are current Good Boys/Girls today
	*[]bool // Compiler Error: syntax error: unexptected [, expecting name
}

func main() {
	// Notice how the error is more of a syntax violation, than a specific error about embedding slices?
}
