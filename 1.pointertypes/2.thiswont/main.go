package main

// RadarEars is a struct
type RadarEars struct{}

// PointerRadarEars is a struct which is a pointer of RadarEars
type PointerRadarEars *RadarEars

// Corgi embeds PointerRadarEars... or does it?
type Corgi struct {
	PointerRadarEars // Compiler Error: embedded type cannot be a pointer
}

func main() {

}
