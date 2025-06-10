package service

import (
	"errors"
	"fmt"
	"roomreserver/entity"
	"roomreserver/utils"
)

type ReservationService struct {
	Reservas []entity.Reservation
}

func (r *ReservationService) CriarReserva(novaReserva entity.Reservation) error {
	listaDeSalasPermitidas := []string{"A", "B", "C"}
	valido := false
	conflito := false

	for _, sala := range listaDeSalasPermitidas {
		if novaReserva.NomeSala == sala {
			valido = true
			break
		}
	}

	if !valido {
		return errors.New("Sala inválida, tente novamente!")
	}

	if !utils.HoraValida(novaReserva.HoraInicio, novaReserva.HoraFim) {
		conflito = true
		return errors.New("Horário inválido: hora de entrada deve ser menor que a de saída.")
	}

	for _, reserva := range r.Reservas {
		if novaReserva.NomeSala == reserva.NomeSala &&
			utils.ConflitoDeTempo(novaReserva.HoraInicio, reserva.HoraFim, reserva.HoraInicio, novaReserva.HoraFim) {
			return errors.New("Horário não está disponível, tente novamente!")
		}
	}

	if !conflito {
		fmt.Println("Horário disponível.")
	}

	r.Reservas = append(r.Reservas, novaReserva)
	return nil
}

func (r *ReservationService) ListarReservasPorSala() []entity.Reservation {
	return r.Reservas
}

func (r *ReservationService) VerificarDisponibilidade() {
	listaDeSalasPermitidas := []string{"A", "B", "C"}
	fmt.Println("As salas disponíveis são ->", listaDeSalasPermitidas)
	fmt.Println("")
}
