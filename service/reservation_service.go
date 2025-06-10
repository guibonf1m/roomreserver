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

	// Verifica se a sala da nova reserva é uma das salas permitidas.
	for _, sala := range listaDeSalasPermitidas {
		if novaReserva.NomeSala == sala {
			valido = true
			break // Encontra a sala válida e não precisa continuar verificando outras.
		}
	}

	// Caso a sala não esteja na lista de permitidas, avisa e encerra a operação.
	if !valido {
		return errors.New("A sala é inválida, tente novamente!")
	}

	// Percorre todas as reservas existentes para verificar conflito de horário.
	for _, reserva := range r.Reservas {
		if novaReserva.NomeSala == reserva.NomeSala { // Verifica apenas reservas na mesma sala.
			// Usa a função utilitária para confirmar se há sobreposição de horários.
			if utils.ConflitoDeTempo(novaReserva.HoraInicio, reserva.HoraFim, reserva.HoraInicio, novaReserva.HoraFim) {
				conflito = true                                                    // Marca como verdadeiro ao encontrar conflito.
				return errors.New("Horário não está disponível, tente novamente!") // Encerra a operação se houver conflito.
			}
		}
	}

	// Se não houve nenhum conflito, avisa que o horário está disponível.
	if !conflito {
		fmt.Println("Horário disponível.")
	}

	// Adiciona a nova reserva ao slice de reservas.
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
