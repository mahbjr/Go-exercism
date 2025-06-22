package perfect

import "errors"

// Classification representa o tipo de número
type Classification int

const (
	// ClassificationDeficient indica que a soma dos divisores é menor que o número
	ClassificationDeficient Classification = iota
	// ClassificationPerfect indica que a soma dos divisores é igual ao número
	ClassificationPerfect
	// ClassificationAbundant indica que a soma dos divisores é maior que o número
	ClassificationAbundant
)

// ErrOnlyPositive é o erro retornado para números não positivos
var ErrOnlyPositive = errors.New("Classification is only defined for positive integers")

// Classify determina se um número é perfeito, abundante ou deficiente
func Classify(n int64) (Classification, error) {
	// Verificar se o número é positivo
	if n <= 0 {
		return 0, ErrOnlyPositive
	}

	// Caso especial: 1 não tem divisores próprios além de si mesmo
	if n == 1 {
		return ClassificationDeficient, nil
	}

	// Encontrar a soma dos divisores (exceto o próprio número)
	var sumOfFactors int64 = 1 // 1 é sempre um divisor

	// Só precisamos verificar até a raiz quadrada do número
	for i := int64(2); i*i <= n; i++ {
		if n%i == 0 {
			// i é um divisor
			sumOfFactors += i

			// Se i não é a raiz quadrada, então n/i também é um divisor
			if i != n/i {
				sumOfFactors += n / i
			}
		}
	}

	// Classificar o número com base na soma dos divisores
	if sumOfFactors == n {
		return ClassificationPerfect, nil
	} else if sumOfFactors > n {
		return ClassificationAbundant, nil
	} else {
		return ClassificationDeficient, nil
	}
}
