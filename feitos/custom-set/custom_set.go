package stringset

import (
	"sort"
	"strings"
)

type Set map[string]bool

// New returns a new empty Set
func New() Set {
	return make(Set)
}

// NewFromSlice returns a new Set populated from a slice of strings
func NewFromSlice(slice []string) Set {
	set := New()
	if slice != nil {
		for _, s := range slice {
			set.Add(s)
		}
	}
	return set
}

// Add adds an element to the Set
func (s Set) Add(element string) {
	s[element] = true
}

// Delete removes an element from the Set
func (s Set) Delete(element string) {
	delete(s, element)
}

// Has returns true if the element is in the Set
func (s Set) Has(element string) bool {
	return s[element]
}

// IsEmpty returns true if the Set contains no elements
func (s Set) IsEmpty() bool {
	return len(s) == 0
}

// Slice returns a slice containing all elements of the Set
func (s Set) Slice() []string {
	slice := make([]string, 0, len(s))
	for elem := range s {
		slice = append(slice, elem)
	}
	return slice
}

// String returns a string representation of the Set
func (s Set) String() string {
	elements := s.Slice()
	sort.Strings(elements) // Para ter uma ordem consistente nos testes
	
	var builder strings.Builder
	builder.WriteString("{")
	for i, elem := range elements {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(`"` + elem + `"`)
	}
	builder.WriteString("}")
	return builder.String()
}

// Equal returns true if two Sets contain exactly the same elements
func Equal(s1, s2 Set) bool {
	if len(s1) != len(s2) {
		return false
	}
	for elem := range s1 {
		if !s2.Has(elem) {
			return false
		}
	}
	return true
}

// Subset returns true if all elements of s1 are in s2
func Subset(s1, s2 Set) bool {
	for elem := range s1 {
		if !s2.Has(elem) {
			return false
		}
	}
	return true
}

// Disjoint returns true if s1 and s2 have no elements in common
func Disjoint(s1, s2 Set) bool {
	for elem := range s1 {
		if s2.Has(elem) {
			return false
		}
	}
	return true
}

// Intersection returns a new Set with elements common to both s1 and s2
func Intersection(s1, s2 Set) Set {
	result := New()
	for elem := range s1 {
		if s2.Has(elem) {
			result.Add(elem)
		}
	}
	return result
}

// Difference returns a new Set with elements in s1 that are not in s2
func Difference(s1, s2 Set) Set {
	result := New()
	for elem := range s1 {
		if !s2.Has(elem) {
			result.Add(elem)
		}
	}
	return result
}

// Union returns a new Set with all elements from both s1 and s2
func Union(s1, s2 Set) Set {
	result := New()
	for elem := range s1 {
		result.Add(elem)
	}
	for elem := range s2 {
		result.Add(elem)
	}
	return result
}