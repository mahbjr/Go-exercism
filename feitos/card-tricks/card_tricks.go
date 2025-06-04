package cards

// FavoriteCards returns a slice with the cards 2, 6, and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
}

// GetItem retrieves an item from a slice at the given index.
// If the index is out of range, returns -1.
func GetItem(slice []int, index int) int {
	if index < 0 || index >= len(slice) {
		return -1
	}
	return slice[index]
}

// SetItem modifies a slice to replace the value at a given index.
// If the index is out of range, append the value to the slice.
func SetItem(slice []int, index int, value int) []int {
	if index < 0 || index >= len(slice) {
		return append(slice, value)
	}

	slice[index] = value
	return slice
}

// PrependItems adds items to the beginning of a slice.
func PrependItems(slice []int, values ...int) []int {
	if len(values) == 0 {
		return slice
	}

	// Create a new slice with enough space for all items
	result := make([]int, len(values)+len(slice))

	// Copy the values to the beginning
	copy(result, values)

	// Copy the original slice after the values
	copy(result[len(values):], slice)

	return result
}

// RemoveItem removes an item from a slice at the given index.
// If the index is out of range, return the original slice unchanged.
func RemoveItem(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}

	// Create a new slice excluding the item at the given index
	return append(slice[:index], slice[index+1:]...)
}
