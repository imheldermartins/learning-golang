package examples

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	Role      string `json:"role,omitempty"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
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

func DisplaySerialization() {
	user := User{
		ID:        1,
		Name:      "Helder Martins",
		Email:     "heldi@gmail.com",
		Role:      "superuser",
		Password:  "senha123",
		CreatedAt: "YYYY-MM-DDT00:00:00.000z",
		UpdatedAt: "YYYY-MM-DDT00:00:00.000z",
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
