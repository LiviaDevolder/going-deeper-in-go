# 🧠 Anotações

## ⚙️ Comandos úteis de Go (guia rápido)

### 📦 Módulo e dependências

```bash
go mod init nome-do-modulo       # Inicializa um novo módulo Go
go mod tidy                      # Remove dependências não utilizadas e adiciona as necessárias
go mod download                  # Baixa todas as dependências do módulo
```

### ▶️ Execução

```bash
go run main.go                   # Executa um único arquivo
go run .                         # Executa o pacote atual (todos os arquivos .go)
```

### 🧪 Testes

```bash
go test                          # Roda testes do pacote atual
go test ./...                    # Roda testes de todos os pacotes do módulo
go test -v                       # Testes com saída detalhada
```

### 🔍 Instalação de pacotes

```bash
go get github.com/user/package  # Instala/atualiza um pacote
```

### 🛠️ Build

```bash
go build                         # Compila o pacote atual
go build -o appname              # Compila e define nome do executável
```

### 🧹 Limpeza

```bash
go clean                         # Remove arquivos gerados pela compilação
```

### 📖 Ajuda

```bash
go help                          # Exibe ajuda geral
go help comando                  # Exibe ajuda específica de um comando
```

## 🧠 Tipos básicos de variáveis em Go

### 📦 1. Números

#### 🔢 Inteiros

- `int`, `int8`, `int16`, `int32`, `int64`
- `uint`, `uint8` (alias de `byte`), `uint16`, `uint32`, `uint64`

```go
var a int = 10
var b int64 = 10000000000
var c uint8 = 255
```

#### 🔣 Ponto flutuante

- `float32`, `float64`

```go
var x float32 = 3.14
var y float64 = 2.718281828
```

#### 🧮 Complexos

- `complex64`, `complex128`

```go
var z complex64 = complex(1, 2) // 1 + 2i
```

### 🔤 2. Texto

- `string` – sequência de caracteres UTF-8
- `byte` – alias para `uint8`, usado para representar um caractere ou byte individual
- `rune` – alias para `int32`, representa um caractere Unicode

```go
var s string = "Olá, mundo"
var ch byte = 'a'
var ru rune = 'ç'
```

### ✅ 3. Booleano

- `bool` – valores `true` ou `false`

```go
var ativo bool = true
```

## 🕳️ 4. Valor zero (default)

Quando declaradas sem valor, variáveis em Go assumem um **valor zero**:

| Tipo                                                   | Valor zero |
| ------------------------------------------------------ | ---------- |
| `int`                                                  | `0`        |
| `float`                                                | `0.0`      |
| `bool`                                                 | `false`    |
| `string`                                               | `""`       |
| ponteiros, slices, maps, interfaces, channels, funções | `nil`      |

```go
var a int      // 0
var b string   // ""
var c bool     // false
```

### 💡 Dica: Inferência de tipo

Go pode inferir o tipo a partir do valor:

```go
idade := 30         // int
nome := "João"      // string
ativo := true       // bool
```

## 🧪 Testes

### ▶️ Rodar testes

```bash
go test
```

- Executa todos os testes no pacote atual.

```bash
go test ./...
```

- Executa testes em todos os subdiretórios.

---

### 🧪 Rodar testes com detalhes

```bash
go test -v
```

- Mostra a execução de cada teste (modo verbose).

```bash
go test -run NomeDoTeste
```

- Executa apenas o teste que contém `NomeDoTeste`.

---

### 📈 Verificar cobertura

```bash
go test -cover
```

- Mostra a porcentagem de cobertura no terminal.

```bash
go test -coverprofile=cobertura.out
```

- Gera um arquivo com os dados de cobertura.

```bash
go tool cover -html=cobertura.out
```

- Abre um relatório visual da cobertura no navegador.

---

### 🧪 Testes com benchmarks

```bash
go test -bench=.
```

- Executa benchmarks do pacote atual.

---

### 🧪 Criar arquivo de teste

- Nome do arquivo: `nome_arquivo_test.go`
- Estrutura básica:

```go
func TestNome(t *testing.T) {
    resultado := Soma(2, 3)
    esperado := 5

    if resultado != esperado {
        t.Errorf("esperado %d, obtido %d", esperado, resultado)
    }
}
```

---

### 🧹 Limpar dependências não utilizadas

```bash
go mod tidy
```
