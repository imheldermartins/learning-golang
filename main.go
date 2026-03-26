/*
	1. Verbos Genéricos (Qualquer tipo)
		* %v: O valor no formato padrão.
		* %+v: Para structs, inclui os nomes dos campos.
		* %#v: Representação do valor em sintaxe Go.
		* %T: Exibe o tipo Go do valor (ex: int, string).
		* %%: Escreve um sinal de porcentagem literal.
	2. Inteiros
		* %d: Base 10 (decimal).
		* %b: Base 2 (binário).
		* %x / %X: Base 16 (hexadecimal), minúsculo ou maiúsculo.
		* %o: Base 8 (octal).
		* %c: O caractere (rune) correspondente ao código.
	3. Ponto Flutuante (Floats)
		* %f: Decimal sem expoente (ex: 123.456000).
		* %.2f: Decimal com precisão de 2 casas.
		* %e / %E: Notação científica.
		* %g: Escolhe %e ou %f automaticamente para o que for mais curto.
	4. Strings e Booleanos
		* %s: String ou fatia de bytes (sem aspas).
		* %q: String entre aspas duplas, segura para código Go.
		* %t: Booleano (true ou false).
	5. Ponteiros
		* %p: Endereço de memória em hexadecimal com prefixo 0x
*/

package main

import (
	"fmt"
)

/*
Simple function to add two integers
*/
// func sum(a int, b int) int {
// 	return a + b
// }

// func divide(a float64, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("cannot divide by zero.")
// 	}
// 	return a / b, nil
// }

type User struct {
	Name string
}

func (u User) Greet() string {
	return fmt.Sprintf("Hello, %s!", u.Name)
}

func (u *User) setName(newName string) {
	u.Name = newName
}

func main() {
	// const pi float64 = 3.14159

	// var a int = 2
	// var b int = 3

	// result := a + b

	// price := 19.99

	// var isAdmin bool = true

	// var status string

	// if isAdmin {
	// 	status = "Yes"
	// } else {
	// 	status = "No"
	// }

	// fmt.Printf("Is the user an admin? %s\n", status)

	// fmt.Printf("The price is $%.2f\n", price)

	// fmt.Printf("%d + %d = %d\n", a, b, result)

	// first_name := "Helder"
	// last_name := "Martins"

	// full_name := first_name + " " + last_name

	// fmt.Printf("Hello, %s!\n", full_name)

	// const limit int = 10

	// if result := sum(6, 5); result >= limit {
	// 	fmt.Printf("The result is %d, which is greater than or equal to %d.\n", result, limit)
	// }

	// for i := 0;i <= 10;i++ {
	// 	fmt.Printf("> %d\n", i)
	// }

	// index := 0

	// for true {
	// 	fmt.Printf("> %d\n", index)
	// 	index++
	// 	if index > 10 {
	// 		break
	// 	}
	// }

	// array := []int{10, 20, 30}

	// fmt.Println("O indíce 0 é: ", array[0])
	// for index, value := range array {
	// 	fmt.Println(index, value)
	// }

	// fmt.Println("================================")
	// array = append(array, 40)

	// for index, value := range array {
	// 	fmt.Println(index, value)
	// }
	// fmt.Println("================================")

	// array = array[:len(array)-1]
	// for index, value := range array {
	// 	fmt.Println(index, value)
	// }

	// var a, b float64 = 3, 2
	// var result float64 = math.Pow(a, b)

	// fmt.Printf("%.2f^%.2f = %.2f\n", a, b, result)
	// fmt.Printf("Sqrt(%d) = %d\n", int(result), int(math.Sqrt(result)))

	user := User{Name: "Helder"}

	fmt.Println(user.Greet())
	user.setName("Helder Martins")
	fmt.Println(user.Greet())
}
