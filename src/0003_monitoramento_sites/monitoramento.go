package main

import (
	"fmt"
)

func main() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair")

	var comando int
	fmt.Scan(&comando)
	print("--> ", comando, " || ")

	switch comando {
	case 1:
		fmt.Println("Monitorando...")

	case 2:
		fmt.Println("Recuperando logs...")

	case 0:
		fmt.Println("Aplicação Encerrada")

	default:
		fmt.Println("Entrada Inválida!!!")
	}
}
