package main

// RadarEars is a struct
type RadarEars struct {
	CutenessLevel int
}

// Corgi embeds RadarEars
type Corgi struct {
	RadarEars
	Name          string
	CutenessLevel int
}

func main() {
	corg := Corgi{Name: "Buster Brown"}
	corg.CutenessLevel = 100 // CutenessLevel of Corgi is now set, but RadarEars is not!

	corg2 := Corgi{Name: "Aria"}
	corg2.RadarEars.CutenessLevel = 100 // CutenessLevel of RadarEars are set, but Corgi is unchanged!
}
