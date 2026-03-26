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

## Executar com Docker

```sh
docker build -t learning . && docker run learning:latest
```