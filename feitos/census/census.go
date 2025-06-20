// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	// Check if name is provided
	if r.Name == "" {
		return false
	}

	// Check if address is provided and not empty
	if r.Address == nil || len(r.Address) == 0 {
		return false
	}

	// Check if the address has a street field with a non-empty value
	street, hasStreet := r.Address["street"]
	if !hasStreet || street == "" {
		return false
	}

	// Age is optional, so no need to check it
	return true
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Name = ""
	r.Age = 0
	r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	count := 0

	for _, resident := range residents {
		if resident.HasRequiredInfo() {
			count++
		}
	}

	return count
}
