package utils

// ConflitoDeTempo verifica se há sobreposição entre dois intervalos de tempo
func ConflitoDeTempo(start1, end1, start2, end2 int) bool {
	// Retorna true se os intervalos se sobrepõem
	return start1 < end2 && end1 > start2
}
