# Aprendendo Golang

## Tópicos Cobertos

### 1. Formatação de Strings
- **Genéricos**: `%v` (valor), `%+v` (com campos), `%#v` (sintaxe Go), `%T` (tipo)
- **Inteiros**: `%d` (decimal), `%b` (binário), `%x/%X` (hexadecimal), `%o` (octal), `%c` (rune)
- **Floats**: `%f` (decimal), `%.2f` (precisão), `%e/%E` (científica), `%g` (automático)
- **Strings/Booleanos**: `%s` (string), `%q` (com aspas), `%t` (booleano)
- **Ponteiros**: `%p` (endereço de memória)

### 2. Funções
- Declaração com múltiplos parâmetros e retornos
- Padrão de erro: retornar `(resultado, error)`
- Exemplo:
```go
func divide(a float64, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("não pode dividir por zero")
    }
    return a / b, nil
}
```

### 3. Controle de Fluxo
- `if/else` com condições
- Loops `for`: tradicional, infinito com `break`, e iteração com `range`

### 4. Estruturas de Dados
- Arrays e slices: criação, indexação, iteração com `range`
- Operações: `append()` e slicing `[:len(array)-1]`

### 5. Programação Orientada a Objetos
- **Structs**: Definição de tipos customizados
- **Métodos com Receivers**: `(u User)` para leitura, `(u *User)` para modificação
```go
type User struct {
    Name      string
    Email     string
    Admin     bool
    CreatedAt string
}

func (u User) getEmail() (string, error) {
    if u.Email == "" {
        return "", fmt.Errorf("Email não definido")
    }
    return u.Email, nil
}

func (u *User) setEmail(newEmail string) {
    u.Email = newEmail
}
```

### 6. Tratamento de Erros
- Retornar erros de funções e métodos
- Validar valores antes de usar
- Padrão: `if err != nil { return }`

### 7. Serialização JSON
- **Tags JSON**: Mapeiam campos da struct para nomes JSON
  - `json:"fieldname"` - Serializa o campo com o nome especificado
  - `json:"-"` - Omite o campo completamente do JSON (nunca aparece)
  - `json:"fieldname,omitempty"` - Omite o campo se estiver vazio
- **Serialização**: `json.Marshal()` converte struct para bytes JSON
- **Desserialização**: `json.Unmarshal()` converte JSON para struct
- Exemplo:
```go
type User struct {
    ID        int    `json:"id"`
    Name      string `json:"name"`
    Email     string `json:"email"`
    Password  string `json:"-"`                // Nunca aparece no JSON
    Role      string `json:"role,omitempty"`   // Omitido se vazio
    CreatedAt string `json:"createdAt"`
    UpdatedAt string `json:"updatedAt"`
}

user := User{
    ID:        1,
    Name:      "Helder Martins",
    Email:     "heldi@gmail.com",
    Role:      "superuser",
    Password:  "senha123",                     // Será ignorado
    CreatedAt: "YYYY-MM-DDT00:00:00.000z",
    UpdatedAt: "YYYY-MM-DDT00:00:00.000z",
}

// Serializar para JSON
data, _ := json.Marshal(user)
fmt.Println(string(data))
// Output: {"id":1,"name":"Helder Martins","email":"heldi@gmail.com","role":"superuser","createdAt":"YYYY-MM-DDT00:00:00.000z","updatedAt":"YYYY-MM-DDT00:00:00.000z"}
// Note: Password não aparece (json:"-") e Role aparece (não está vazio)
```

## Executar com Docker

```sh
docker build -t learning . && docker run learning:latest
```