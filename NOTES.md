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
