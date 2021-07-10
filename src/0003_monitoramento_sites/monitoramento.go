package main

import (
	"fmt"
)

func main() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair")

	var comando int
	fmt.Scan(&comando)

	println("-->", comando)
}
