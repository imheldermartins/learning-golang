package examples

/*
import "learning/examples" // imports the examples package

func main() {
	examples.DisplaySerialization()
}
*/

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	jsonData := []byte(`{"name": "Alice", "age": 30}`)

	var person Person

	err := json.Unmarshal(jsonData, &person)

	if err != nil {
		fmt.Println("Occured an error: ", err)
		return
	}

	fmt.Println("Name: ", person.Name)
	fmt.Println("Age: ", person.Age)
}
