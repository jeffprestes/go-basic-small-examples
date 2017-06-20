package main

import (
	"fmt"
	"strings"
)

func main() {
	priNome, segNome, terNome, quaNome, quiNome := nomeDoJeff("Jefferson Prestes de Oliveira Barbosa")
	fmt.Printf("1o.: %s - 2o.: %s - 3o.: %s - 4o.: %s - 5o.: %s\n", priNome, segNome, terNome, quaNome, quiNome)
}

func nomeDoJeff(nomeCompleto string) (string, string, string, string, string) {
	arrayNomes := strings.Split(nomeCompleto, " ")
	return arrayNomes[0], arrayNomes[1], arrayNomes[2], arrayNomes[3], arrayNomes[4]
}
