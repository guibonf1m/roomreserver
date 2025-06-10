package utils

import "math/rand"

// ConflitoDeTempo verifica se há sobreposição entre dois intervalos de tempo
func ConflitoDeTempo(start1, end1, start2, end2 int) bool {
	// Retorna true se os intervalos se sobrepõem
	return start1 < end2 && end1 > start2
}

func GerarIdAleatorio() string {
	const letras = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	// Aqui estão todos os símbolos possíveis para o ID: letras minúsculas, maiúsculas e números

	n := 8                // O ID terá 8 caracteres
	id := make([]byte, n) // Cria uma “caixa” para guardar 8 símbolos, cada um é um byte

	for i := 0; i < n; i++ { // Para cada posição (de 0 até 7) do ID...
		id[i] = letras[rand.Intn(len(letras))]
		// - rand.Intn(len(letras)) sorteia um número de 0 até a quantidade de letras e números disponíveis.
		// - letras[...] pega o caractere nessa posição sorteada da string 'letras'.
		// - id[i] = ... coloca esse caractere sorteado na posição i do slice 'id', montando o ID final passo a passo.

	}

	return string(id) // Junta todos os bytes gerados e transforma em texto antes de retornar
}
