package utils

import "math/rand"

func ConflitoDeTempo(start1, end1, start2, end2 int) bool {

	return start1 < end2 && end1 > start2
}

func HoraValida(horaentrada, horasaida int) bool {

	return horaentrada < horasaida
}

func GerarIdAleatorio() string {
	const letras = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	n := 8
	id := make([]byte, n)

	for i := 0; i < n; i++ {
		id[i] = letras[rand.Intn(len(letras))]
	}

	return string(id)
}
