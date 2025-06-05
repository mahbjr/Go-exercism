package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	// Check if the unit is valid
	quantity, ok := units[unit]
	if !ok {
		return false
	}

	// If item exists, add to the current quantity
	if currentQty, exists := bill[item]; exists {
		bill[item] = currentQty + quantity
	} else {
		bill[item] = quantity
	}

	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	// Check if the item is in the bill
	currentQty, itemExists := bill[item]
	if !itemExists {
		return false
	}

	// Check if the unit is valid
	unitQty, unitExists := units[unit]
	if !unitExists {
		return false
	}

	// Calculate the new quantity
	newQty := currentQty - unitQty

	// Check if removing would result in negative quantity
	if newQty < 0 {
		return false
	}

	// If new quantity is 0, remove the item from the bill
	if newQty == 0 {
		delete(bill, item)
	} else {
		bill[item] = newQty
	}

	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, exists := bill[item]
	return quantity, exists
}
