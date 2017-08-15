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

func main() {
	// Pretend like we got this from a database!
	u := User{
		Name:           "Josephina Userface",
		HashedPassword: []byte{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		DarkSecret:     "Hates tea",
	}
	// Now let's put it into a protected view model, so no secrets leak!
	uv := &UserView{
		User: u,
	}
	// Let's marshal both objects, to compare them
	uj, _ := json.MarshalIndent(u, "", "\t")
	uvj, _ := json.MarshalIndent(uv, "", "\t")
	fmt.Println(string(uj))
	fmt.Println(string(uvj))

	// Now, what happens if we alter the UserView?
	uv.DarkSecret = "b"
	uv.Name = "Doof"
	uj, _ = json.MarshalIndent(u, "", "\t")
	uvj, _ = json.MarshalIndent(uv, "", "\t")
	fmt.Println(string(uj))
	fmt.Println(string(uvj))

	// Even if embedded type is a pointer, the underlying user is never directly altered
	// All the important data is represented via the viewmodels state!
}
