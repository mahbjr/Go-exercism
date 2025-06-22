package letter

import "sync"

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
	frequencies := FreqMap{}
	for _, r := range text {
		frequencies[r]++
	}
	return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given strings,
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
	// Criar o resultado final
	result := FreqMap{}

	// Se não houver textos para processar
	if len(texts) == 0 {
		return result
	}

	// Canal para receber os resultados de cada goroutine
	freqChan := make(chan FreqMap, len(texts))

	// Inicia uma goroutine para cada texto
	var wg sync.WaitGroup
	for _, text := range texts {
		wg.Add(1)
		// Captura a variável text no escopo da função anônima
		go func(t string) {
			defer wg.Done()
			freqChan <- Frequency(t)
		}(text)
	}

	// Goroutine para fechar o canal depois que todas as tarefas terminarem
	go func() {
		wg.Wait()
		close(freqChan)
	}()

	// Mescla os resultados de todas as goroutines
	for freqMap := range freqChan {
		for r, count := range freqMap {
			result[r] += count
		}
	}

	return result
}
