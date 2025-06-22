package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Define the Robot type here.
type Robot struct {
	name string
}

// Mantém registro de nomes já utilizados
var (
	usedNames    = make(map[string]bool)
	namesMutex   sync.Mutex
	initialized  bool
)

// Inicializa o gerador de números aleatórios
func init() {
	rand.Seed(time.Now().UnixNano())
	initialized = true
}

// Name retorna o nome do robô, gerando um novo se necessário
func (r *Robot) Name() (string, error) {
	// Se o robô já tem um nome, retorna o nome existente
	if r.name != "" {
		return r.name, nil
	}

	// Tenta gerar um novo nome único
	namesMutex.Lock()
	defer namesMutex.Unlock()

	// Verifica se todos os nomes possíveis já foram usados
	if len(usedNames) >= 26*26*10*10*10 {
		return "", errors.New("namespace is exhausted")
	}

	// Tenta gerar um nome único até conseguir ou até 100 tentativas
	for attempt := 0; attempt < 100; attempt++ {
		// Gera um novo nome no formato "AA###"
		name := generateName()

		// Verifica se o nome já foi usado
		if !usedNames[name] {
			usedNames[name] = true
			r.name = name
			return name, nil
		}
	}

	// Se não conseguiu gerar um nome único após muitas tentativas,
	// faz uma verificação sequencial (menos eficiente, mas mais garantido)
	return generateSequential()
}

// Reset limpa o nome atual do robô, para que um novo seja gerado na próxima chamada a Name()
func (r *Robot) Reset() {
	r.name = ""
}

// generateName gera um nome aleatório no formato "AA###"
func generateName() string {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"

	letter1 := letters[rand.Intn(26)]
	letter2 := letters[rand.Intn(26)]
	digit1 := digits[rand.Intn(10)]
	digit2 := digits[rand.Intn(10)]
	digit3 := digits[rand.Intn(10)]

	return fmt.Sprintf("%c%c%c%c%c", letter1, letter2, digit1, digit2, digit3)
}

// generateSequential tenta todos os nomes possíveis sequencialmente até encontrar um não usado
func generateSequential() (string, error) {
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Tenta todas as combinações possíveis
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			for k := 0; k < 10; k++ {
				for l := 0; l < 10; l++ {
					for m := 0; m < 10; m++ {
						name := fmt.Sprintf("%c%c%d%d%d",
							letters[i], letters[j], k, l, m)
						if !usedNames[name] {
							usedNames[name] = true
							return name, nil
						}
					}
				}
			}
		}
	}

	return "", errors.New("namespace is exhausted")
}
