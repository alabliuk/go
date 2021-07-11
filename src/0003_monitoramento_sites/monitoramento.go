package main

import (
	"fmt"
	"os"
)

func main() {

	comando := menuMain()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")

	case 2:
		fmt.Println("Recuperando logs...")

	case 0:
		fmt.Println("Aplicação Encerrada")
		os.Exit(0)

	default:
		fmt.Println("Entrada Inválida!!!")
		os.Exit(-1)
	}
}

func menuMain() int {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair")

	var comando int
	fmt.Scan(&comando)
	print("--> ", comando, " || ")
	return comando
}
