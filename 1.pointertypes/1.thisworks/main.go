package main

// RadarEars is a struct
type RadarEars struct{}

// Corgi embeds *RadarEars
type Corgi struct {
	*RadarEars // works fine, because it's a pointer to a type, not a "pointer-type"
}

func main() {

}
