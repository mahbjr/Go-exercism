package allyourbase

import (
	"errors"
)

// ConvertToBase converte um número representado por dígitos em uma base de entrada para dígitos em uma base de saída
func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	// Validação das bases
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	// Tratamento de casos especiais: lista vazia ou múltiplos zeros
	if len(inputDigits) == 0 {
		return []int{0}, nil
	}

	// Remove leading zeros
	digits := removeLeadingZeros(inputDigits)
	if len(digits) == 1 && digits[0] == 0 {
		return []int{0}, nil
	}

	// Validação dos dígitos
	for _, d := range digits {
		if d < 0 || d >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
	}

	// Converte para base 10 (decimal)
	decimalValue := 0
	for _, digit := range digits {
		decimalValue = decimalValue*inputBase + digit
	}

	// Caso especial: valor zero
	if decimalValue == 0 {
		return []int{0}, nil
	}

	// Converte da base 10 para a base de saída
	return convertFromDecimal(decimalValue, outputBase), nil
}

// removeLeadingZeros remove zeros à esquerda do slice de dígitos
func removeLeadingZeros(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}

	i := 0
	for i < len(digits) && digits[i] == 0 {
		i++
	}

	if i == len(digits) {
		return []int{0}
	}

	return digits[i:]
}

// convertFromDecimal converte um número decimal para outra base
func convertFromDecimal(number, base int) []int {
	if number == 0 {
		return []int{0}
	}

	var digits []int
	for number > 0 {
		digits = append(digits, number%base)
		number = number / base
	}

	// Inverte os dígitos
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	return digits
}