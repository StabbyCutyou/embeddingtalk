package lamb

import "fmt"

// Treat is a tasty lamb treat for dogs to enjoy
type Treat struct {
	size        int64
	SourcedFrom string
}

// Description will say hellos
func (t *Treat) Description() {
	fmt.Printf("Hi, i'm a size %d Lamb Treat\n", t.size)
}
