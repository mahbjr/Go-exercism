package piglatin

import (
	"strings"
)

// Sentence traduz uma frase inteira para Pig Latin
func Sentence(input string) string {
	words := strings.Fields(input)
	result := make([]string, len(words))

	for i, word := range words {
		result[i] = translateWord(word)
	}

	return strings.Join(result, " ")
}

// translateWord traduz uma única palavra para Pig Latin
func translateWord(word string) string {
	// Regra especial para palavras que começam com "xr" ou "yt"
	if strings.HasPrefix(word, "xr") || strings.HasPrefix(word, "yt") {
		return word + "ay"
	}

	// Palavras que começam com vogal
	if isVowel(word[0]) {
		return word + "ay"
	}

	// Encontra o índice da primeira vogal ou fim de cluster consonantal
	index := findVowelIndex(word)

	// Caso especial para "qu" após consoante(s)
	if index < len(word)-1 && word[index:index+2] == "qu" {
		index += 2
	}

	// Mover consoantes iniciais para o final e adicionar "ay"
	return word[index:] + word[:index] + "ay"
}

// isVowel verifica se um caractere é uma vogal
func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}

// findVowelIndex encontra o índice da primeira vogal na palavra
// com tratamento especial para "qu", "y", etc.
func findVowelIndex(word string) int {
	for i := 0; i < len(word); i++ {
		// Caso de "qu" - manter junto com consoante anterior
		if i < len(word)-1 && word[i] == 'q' && word[i+1] == 'u' {
			i++
			continue
		}

		// 'y' é tratado como vogal quando não está no início da palavra
		// e quando está após uma consoante
		if word[i] == 'y' && i > 0 {
			return i
		}

		// Vogais regulares
		if isVowel(word[i]) {
			return i
		}
	}

	// Se não encontrou vogal, retorna o comprimento da palavra
	return len(word)
}
