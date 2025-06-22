package allergies

// Mapeamento de alergênios para seus valores binários
var allergenValues = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

// Ordem dos alergênios para garantir uma saída consistente
var allergenOrder = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

// Allergies retorna todos os alergênios aos quais uma pessoa com determinado score é alérgica
func Allergies(score uint) []string {
	var allergies []string

	// Apenas os 8 primeiros bits são relevantes para as alergias
	relevantScore := score & 255

	// Verifica cada alergênio na ordem definida
	for _, allergen := range allergenOrder {
		if isAllergic(allergen, relevantScore) {
			allergies = append(allergies, allergen)
		}
	}

	return allergies
}

// isAllergic é uma função auxiliar que verifica se um bit específico está ativo
func isAllergic(allergen string, score uint) bool {
	allergenValue, exists := allergenValues[allergen]
	if !exists {
		return false
	}

	// Usa um AND binário para verificar se o bit correspondente ao alergênio está ativo no score
	return score&allergenValue != 0
}

// AllergicTo verifica se uma pessoa com determinado score é alérgica a um alergênio específico
func AllergicTo(score uint, allergen string) bool {
	return isAllergic(allergen, score)
}

// List é um alias para Allergies para manter compatibilidade com os testes existentes
func List(score uint) []string {
	return Allergies(score)
}
