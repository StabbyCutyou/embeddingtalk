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
	HashedPassword []byte `json:"hashed_password,omitempty"` // If you leave off omitempty
	DarkSecret     string `json:"dark_secret,omitempty"`     // The field will still show up, albeit with a 0 value
}

// UserTagView is a viewmodel
type UserTagView struct {
	User           `json:"user"` // This subtle difference has big implications
	HashedPassword []byte        `json:"hashed_password,omitempty"`
	DarkSecret     string        `json:"dark_secret,omitempty"`
}

func main() {
	// Pretend like we got this from a database!
	u := User{
		Name:           "Doggletons Q Barkface",
		HashedPassword: []byte{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		DarkSecret:     "Hates birds",
	}
	// Now let's put it into a protected view model, so no secrets leak!
	uv := &UserView{
		User: u,
	}
	// Now let's put it into a protected view model, so no secrets leak!
	utv := &UserTagView{
		User: u,
	}
	// Let's marshall both objects, to compare them
	uvj, _ := json.MarshalIndent(uv, "", "\t")
	utj, _ := json.MarshalIndent(utv, "", "\t")

	fmt.Println(string(uvj)) // Without the tag, the marshaller will ignore the User data
	fmt.Println(string(utj)) // Once you explicitly set the tag, the marshaller will include the embedded type!
}
