package encode

import (
	"strconv"
	"strings"
	"unicode"
)

// RunLengthEncode codifica uma string usando run-length encoding
// Por exemplo: "AABBBCCCC" -> "2A3B4C"
func RunLengthEncode(input string) string {
	if len(input) == 0 {
		return ""
	}

	var result strings.Builder
	count := 1
	current := rune(input[0])

	// Percorre cada caractere a partir do segundo
	for _, char := range input[1:] {
		if char == current {
			// Mesmo caractere, incrementa a contagem
			count++
		} else {
			// Caractere diferente, escreve a contagem anterior e o caractere
			if count > 1 {
				result.WriteString(strconv.Itoa(count))
			}
			result.WriteRune(current)

			// Reinicia a contagem com o novo caractere
			count = 1
			current = char
		}
	}

	// Processa o último grupo de caracteres
	if count > 1 {
		result.WriteString(strconv.Itoa(count))
	}
	result.WriteRune(current)

	return result.String()
}

// RunLengthDecode decodifica uma string que foi codificada usando run-length encoding
// Por exemplo: "2A3B4C" -> "AABBBCCCC"
func RunLengthDecode(input string) string {
	if len(input) == 0 {
		return ""
	}

	var result strings.Builder
	var countStr strings.Builder

	for _, char := range input {
		if unicode.IsDigit(char) {
			// Acumula os dígitos do contador
			countStr.WriteRune(char)
		} else {
			// Obtém o contador (ou 1 se não houver contador)
			count := 1
			if countStr.Len() > 0 {
				count, _ = strconv.Atoi(countStr.String())
				countStr.Reset()
			}

			// Repete o caractere o número de vezes especificado
			for i := 0; i < count; i++ {
				result.WriteRune(char)
			}
		}
	}

	return result.String()
}
