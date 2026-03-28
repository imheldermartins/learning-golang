package main

// imports the examples package
import (
	"fmt"
	"learning/api"
)

func main() {
	api.Init() // Initialize the database connection
	// api.GetUsers()
	user, err := api.GetUser(2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User: %+v\n", user)
}
