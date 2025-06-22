package flatten

// Flatten takes a nested array structure and returns a flattened slice
// Null values (nil) are omitted from the result
func Flatten(nested interface{}) []interface{} {
	// Inicializar com uma slice vazia (não nil)
	result := make([]interface{}, 0)

	// Se a entrada for nil, retornar a slice vazia inicializada
	if nested == nil {
		return result
	}

	// Verificar se é uma slice
	if list, ok := nested.([]interface{}); ok {
		for _, item := range list {
			// Pular valores nulos
			if item == nil {
				continue
			}

			// Verificar se o item é uma slice aninhada
			if nestedList, ok := item.([]interface{}); ok {
				// Achatar recursivamente a lista aninhada e adicionar ao resultado
				flattenedSublist := Flatten(nestedList)
				result = append(result, flattenedSublist...)
			} else {
				// Adicionar valores simples diretamente ao resultado
				result = append(result, item)
			}
		}
	} else {
		// Se não for uma slice, adicionar o valor diretamente (se não for nil)
		if nested != nil {
			result = append(result, nested)
		}
	}

	return result
}
