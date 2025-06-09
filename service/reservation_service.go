package service

import (
	"fmt"
	"roomreserver/entity"
	"roomreserver/utils"
)

type ReservationService struct {
	Reservas []entity.Reservation
}

func (r *ReservationService) CriarReserva(novaReserva entity.Reservation) {
	listaDeSalasPermitidas := []string{"A", "B", "C"}
	valido := false

	for _, sala := range listaDeSalasPermitidas {
		if novaReserva.NomeSala == sala {
			valido = true
			fmt.Println("Sala válida", valido)
			break
		}
	}

	if !valido {
		fmt.Println("A sala é inválida, tente novamente!")
	}

	r.CriarReserva()

}

func (r *ReservationService) ListarReservasPorSala() []entity.Reservation {
	return r.Reservas
}

func (r *ReservationService) VerificarDisponibilidade() {

}
