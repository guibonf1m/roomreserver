package main

import (
	"fmt"
	"os"
	"roomreserver/entity"
	"roomreserver/service"
	"roomreserver/utils"
)

func main() {

	var minhasSalas service.ReservationService

	for {
		fmt.Println("1 - Mostrar salas")
		fmt.Println("2 - Criar reservas")
		fmt.Println("3 - Listar reservas")
		fmt.Println("4 - Sair \n")
		fmt.Print("Escolha uma opção: ")

		var opcao string
		fmt.Scanln(&opcao)

		switch opcao {

		case "1":

			var sala string
			var start int
			var end int

			fmt.Println("Digite a sala, (A, B ou C): ")
			fmt.Scanln(&sala)

			fmt.Println("Qual a hora de inicio: ")
			fmt.Scanln(&start)

		case "2":

		case "3":

		case "4":
			fmt.Println("Saindo do Programa...")
			os.Exit(4)

		default:
			fmt.Println("Opção inválida, tente novamente!")
		}

	}

}
