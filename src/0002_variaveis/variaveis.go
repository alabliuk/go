package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Declando os tipos
	var nome string = "Andre"
	var idade int = 21
	var versao float32 = 1.1

	fmt.Println("Ola", nome, ", sua idade é:", idade)
	fmt.Println("Versao", versao)

	// Sem declarar tipos
	var nome2 = "Andre"
	var idade2 = 21
	var versao2 = 1.1

	fmt.Println("Ola", nome2, ", sua idade é:", idade2)
	fmt.Println("Versao", versao2)
	fmt.Println("tipo variavel: ", reflect.TypeOf(versao2))

	// atribuicao de variaveis direta
	nome3 := "Andre"
	idade3 := 21
	versao3 := 1.1

	fmt.Println("Ola", nome3, ", sua idade é:", idade3)
	fmt.Println("Versao", versao3)
}
