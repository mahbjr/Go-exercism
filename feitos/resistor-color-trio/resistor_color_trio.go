package resistorcolortrio

import (
	"fmt"
	"math"
	"strings"
)

// Color representa o mapeamento de cores para valores
var colorValues = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Label calcula o valor do resistor baseado nas cores fornecidas
func Label(colors []string) string {
	if len(colors) < 3 {
		return "invalid input"
	}

	// Pegar os valores das duas primeiras cores
	firstDigit := colorValues[colors[0]]
	secondDigit := colorValues[colors[1]]

	// Calcular o valor base (primeiros dois dígitos)
	baseValue := firstDigit*10 + secondDigit

	// Pegar o valor do multiplicador (terceira cor)
	multiplier := int(math.Pow10(colorValues[colors[2]]))

	// Calcular o valor total em ohms
	totalOhms := baseValue * multiplier

	// Determinar a unidade apropriada
	var unit string
	var adjustedValue float64

	switch {
	case totalOhms < 1000:
		unit = "ohms"
		adjustedValue = float64(totalOhms)
	case totalOhms < 1000000:
		unit = "kiloohms"
		adjustedValue = float64(totalOhms) / 1000
	case totalOhms < 1000000000:
		unit = "megaohms"
		adjustedValue = float64(totalOhms) / 1000000
	default:
		unit = "gigaohms"
		adjustedValue = float64(totalOhms) / 1000000000
	}

	// Formatar o resultado removendo zeros decimais desnecessários
	var result string
	if math.Floor(adjustedValue) == adjustedValue {
		result = fmt.Sprintf("%d %s", int(adjustedValue), unit)
	} else {
		result = fmt.Sprintf("%.1f %s", adjustedValue, unit)
		// Remover ".0" se for um número inteiro
		result = strings.ReplaceAll(result, ".0", "")
	}

	return result
}
