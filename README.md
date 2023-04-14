# Métodos, interfaces e *generics*

## Sobre

Este conjunto de programas serve à demonstração do uso de métodos, interfaces e *generics* na linguagem de programação [Go](https://go.dev).

## Estrutura do repositório

Os arquivos fonte deste repositório, cada um demonstrando um exemplo, estão organizados de acordo com a seguinte estrutura:

```
+─abstractions       ---> Nome do diretório do projeto
  ├─── generics      ---> Diretório com exemplos utilizando generics
       ├─── function ---> Diretório com exemplo de implementação de função com generics
       ├─── struct   ---> Diretório com exemplo de implementação de struct com generics
  ├─── interfaces    ---> Diretório com exemplos utilizando interfaces
       ├─── empty    ---> Diretório com exemplo de implementação com interface vazia
       ├─── figures  ---> Diretório com exemplo de implementação de interface e métodos associados
  ├─── methods       ---> Diretório com exemplo de implementação de métodos
```

## Requisitos

Para a compilação e execução dos programas, os seguintes elementos devem estar devidamente instalados no ambiente de desenvolvimento:

- [Git](https://git-scm.com), como sistema de controle de versões
- [Go](https://go.dev), incluindo compilador, ambiente de execução e outras ferramentas associadas à linguagem de programação Go

## Download, compilação e execução dos programas

No terminal do sistema operacional, insira os seguintes comandos para realizar o *download* da implementação a partir deste repositório Git:

```bash
# Download da implementação a partir do repositório Git
git clone https://github.com/ufrn-golang/abstractions.git
```

Para executar um programa, deve-se primeiro navegar para o respectivo diretório no qual ele se encontra e utilizar o comando `go run` no terminal do sistema operacional. Por exemplo, para executar o programa referente ao arquivo [`generics/function/min-generic`](https://github.com/ufrn-golang/abstractions/tree/master/generics/function/min-generic.go), deve-se utilizar os seguintes comandos:

```bash
# Navegar para o diretório do programa
cd generics/function

# Compilar e executar o programa
go run min-generics.go
```

Caso deseje gerar um executável para o programa em questão, deve-se utilizar o comando `go build` e, em seguida, invocar o arquivo executável gerado. Por exemplo, Por exemplo, para compilar e executar o programa referente ao arquivo [`generics/function/min-generic`](https://github.com/ufrn-golang/abstractions/tree/master/generics/function/min-generic.go), deve-se utilizar os seguintes comandos:

```bash
# Navegar para o diretório do programa
cd generics/function

# Compilação do programa e geração de arquivo executável (com o mesmo nome do arquivo fonte)
go build min-generic.go

# Execução do programa
./min-generic
```
