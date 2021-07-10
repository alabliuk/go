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

	if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Recuperando logs...")
	} else if comando == 0 {
		fmt.Println("Aplicação Encerrada")
	} else {
		fmt.Println("Entrada Inválida!!!")
	}
}
