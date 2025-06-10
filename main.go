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
		fmt.Println("2 - Criar nova reserva")
		fmt.Println("3 - Listar reservas")
		fmt.Println("4 - Sair \n")
		fmt.Print("Escolha uma opção: ")

		var opcao string
		fmt.Scanln(&opcao)
		fmt.Println("")

		switch opcao {

		case "1":

			minhasSalas.VerificarDisponibilidade()

		case "2":

			var SalaRecebida string
			var Start int
			var End int
			var Nome string

			fmt.Print("Qual é o seu nome: ")
			fmt.Scanln(&Nome)

			id := utils.GerarIdAleatorio()
			fmt.Println("O seu id aleatório: ", id)
			fmt.Println("")

			fmt.Print("Qual sala você gostaria de reservar (A, B ou C): ")
			fmt.Scanln(&SalaRecebida)

			fmt.Print("Qual será a hora de entrada: ")
			fmt.Scanln(&Start)

			fmt.Print("Qual será a hora de saída: ")
			fmt.Scanln(&End)

			novaReserva := entity.Reservation{
				ReservaId:    id,
				NomeSala:     SalaRecebida,
				HoraInicio:   Start,
				HoraFim:      End,
				ReservadoPor: Nome,
			}

			err := minhasSalas.CriarReserva(novaReserva)
			if err != nil {
				fmt.Println("Erro", err.Error())
				break
			} else {
				fmt.Println("Reserva criada com sucesso!")
			}

		case "3":
			for _, reservas := range minhasSalas.ListarReservasPorSala() {
				fmt.Printf("Reserva ID: %s, Sala: %s, Start: %d, End: %d, Reservador Por: %s ",
					reservas.ReservaId, reservas.NomeSala, reservas.HoraInicio,
					reservas.HoraFim, reservas.ReservadoPor)
				fmt.Println("")
			}

		case "4":
			fmt.Println("Saindo do Programa...")
			os.Exit(4)

		default:
			fmt.Println("Opção inválida, tente novamente!")
		}
	}
}
