package main

import (
	"encoding/json"
	"fmt"
)

// User is a user
type User struct {
	Name           string `json:"name"`
	HashedPassword []byte `json:"hashed_password"`
	DarkSecret     string `json:"dark_secret"`
}

// UserView is a viewmodel
type UserView struct {
	User
	HashedPassword []byte `json:"hashed_password,omitempty"`
	DarkSecret     string `json:"dark_secret,omitempty"`
}

// UserTagView is a viewmodel
type UserTagView struct {
	User           `json:"user,omitempty"`
	HashedPassword []byte `json:"hashed_password,omitempty"`
	DarkSecret     string `json:"dark_secret,omitempty"`
}

func main() {

	// Now, unmarshaling?
	embedded := `{
		"user": {
			"name": "Doggletons Q Barkfacee",
			"hashed_password": "AAAAAQAAAAAAAAAAAAAAAAA=",
			"dark_secret": "Hates birds"
		}
	}`

	full := `{
		"name": "Doggletons Q Barkface",
		"hashed_password": "AAAAAQAAAAAAAAAAAAAAAAA=",
		"dark_secret": "Hates birds"
	}`

	// There are 4! possible ways this can go down!
	newUFV := &UserView{}
	newUEV := &UserView{}
	newUTFV := &UserTagView{}
	newUETV := &UserTagView{}

	// i like to explicitly ignore errors with _ as a visual note that i'm choosing to do this consciously.
	_ = json.Unmarshal([]byte(full), newUFV)
	_ = json.Unmarshal([]byte(embedded), newUEV)
	_ = json.Unmarshal([]byte(full), newUTFV)
	_ = json.Unmarshal([]byte(embedded), newUETV)

	// What do the raw objects look like?
	// Now, how does promotion behave in this situation? What is the value of the DarkSecret field?

	// These cases all operate as expected - boring, let's move past them
	// fmt.Printf("UserView (Full): \t\t%+v\n", newUFV)                          // Set
	// fmt.Printf("\nUserView (Full) DarkSecret: \t\t%s\n", newUFV.DarkSecret)   // Set
	// fmt.Printf("UserView (Embedded): \t%+v\n", newUEV)                        // Not set!
	// fmt.Printf("UserView (Embedded) DarkSecret: \t\t%s\n", newUEV.DarkSecret) // Not set!

	fmt.Printf("UserTagView (Full): \t\t\t%+v\n", newUTFV)                      // Not set!	<<<
	fmt.Printf("UserTagView (Full) DarkSecret: \t\t%s\n", newUTFV.DarkSecret)   // Set		<<<
	fmt.Printf("UserTagView (Embedded): \t\t%+v\n", newUETV)                    // Set		<<<
	fmt.Printf("UserTagView (Embedded) DarkSecret: \t%s\n", newUETV.DarkSecret) // Not set!	<<<

}
