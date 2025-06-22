package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix representa uma matriz de números inteiros
type Matrix [][]int

// New cria uma nova matriz a partir de uma string
func New(s string) (Matrix, error) {
	// Verificar se a string de entrada está vazia
	if s == "" {
		return nil, errors.New("matriz vazia")
	}

	// Dividir a string em linhas
	rows := strings.Split(s, "\n")

	// Verificar linhas vazias
	for _, row := range rows {
		if strings.TrimSpace(row) == "" {
			return nil, errors.New("linha vazia")
		}
	}

	// Criar matriz
	m := make(Matrix, len(rows))

	// Dividir cada linha em números e preencher a matriz
	var width int

	for i, row := range rows {
		// Separar os números da linha
		fields := strings.Fields(row)

		// Na primeira linha, definir a largura da matriz
		if i == 0 {
			width = len(fields)
		} else if len(fields) != width {
			// Verificar se todas as linhas têm o mesmo número de colunas
			return nil, errors.New("linhas com quantidade desigual de elementos")
		}

		// Criar a linha da matriz
		m[i] = make([]int, width)

		// Converter cada campo para um número inteiro
		for j, field := range fields {
			num, err := strconv.Atoi(field)
			if err != nil {
				return nil, errors.New("valor não é um inteiro")
			}
			m[i][j] = num
		}
	}

	return m, nil
}

// Rows retorna uma cópia das linhas da matriz
func (m Matrix) Rows() [][]int {
	// Se a matriz estiver vazia, retornar slice vazio
	if len(m) == 0 {
		return [][]int{}
	}

	// Criar uma cópia profunda da matriz para evitar modificações externas
	result := make([][]int, len(m))
	for i := range m {
		result[i] = make([]int, len(m[i]))
		copy(result[i], m[i])
	}

	return result
}

// Cols retorna uma cópia das colunas da matriz
func (m Matrix) Cols() [][]int {
	// Se a matriz estiver vazia, retornar slice vazio
	if len(m) == 0 || len(m[0]) == 0 {
		return [][]int{}
	}

	// Determinar o número de colunas
	cols := len(m[0])

	// Criar a matriz de resultado
	result := make([][]int, cols)

	// Para cada coluna
	for j := 0; j < cols; j++ {
		// Criar a coluna com tamanho igual ao número de linhas
		result[j] = make([]int, len(m))

		// Copiar os valores da coluna
		for i := 0; i < len(m); i++ {
			result[j][i] = m[i][j]
		}
	}

	return result
}

// Set define o valor de um elemento específico na matriz
// Retorna false se a posição for inválida, true caso contrário
func (m Matrix) Set(r, c, val int) bool {
	// Verificar se as coordenadas estão dentro dos limites da matriz
	if r < 0 || r >= len(m) || c < 0 || c >= len(m[0]) {
		return false
	}

	// Definir o valor
	m[r][c] = val
	return true
}
