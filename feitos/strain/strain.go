package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

// Keep returns a new collection containing all elements where the predicate function returns true.
// It does not modify the input collection.
func Keep[T any](collection []T, predicate func(T) bool) []T {
	var result []T

	// Handle nil collection case
	if collection == nil {
		return nil
	}

	// Filter elements based on the predicate
	for _, item := range collection {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

// Discard returns a new collection containing all elements where the predicate function returns false.
// It does not modify the input collection.
func Discard[T any](collection []T, predicate func(T) bool) []T {
	var result []T

	// Handle nil collection case
	if collection == nil {
		return nil
	}

	// Filter elements based on the inverse of the predicate
	for _, item := range collection {
		if !predicate(item) {
			result = append(result, item)
		}
	}

	return result
}
