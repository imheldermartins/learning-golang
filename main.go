package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Admin     bool   `json:"admin"`
	CreatedAt string `json:"createdAt"`
}

func (u User) getEmail() (string, error) {
	if u.Email == "" {
		return "", fmt.Errorf("Email is not defined.")
	}
	return u.Email, nil
}

func (u *User) setEmail(newEmail string) {
	u.Email = newEmail
}

func main() {
	user := User{
		Name:      "Helder Martins",
		Email:     "heldi@gmail.com",
		Admin:     true,
		CreatedAt: "YYYY-MM-DDT00:00:00.000z",
	}

	email, err := user.getEmail()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Hello %s!\n", email)

	data, _ := json.Marshal(user)

	fmt.Println(string(data))
}
