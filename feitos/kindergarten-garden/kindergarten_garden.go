package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

// Plant codes and their corresponding names
var plantNames = map[byte]string{
	'V': "violets",
	'R': "radishes",
	'C': "clover",
	'G': "grass",
}

// Garden represents a kindergarten garden with plants assigned to children
type Garden struct {
	plantsByChild map[string][]string
}

// NewGarden creates a new garden with plants assigned to children
func NewGarden(diagram string, children []string) (*Garden, error) {
	// Validate diagram format
	rows := strings.Split(diagram, "\n")
	if len(rows) != 3 || rows[0] != "" {
		return nil, errors.New("invalid diagram format: must start with a newline and have two rows of plants")
	}

	// Get the two rows of plants
	row1, row2 := rows[1], rows[2]

	// Check if rows have the same length
	if len(row1) != len(row2) {
		return nil, errors.New("invalid diagram: rows must have the same length")
	}

	// Check if the number of cups is even
	if len(row1)%2 != 0 {
		return nil, errors.New("invalid diagram: odd number of cups")
	}

	// Check if all plant codes are valid
	for i := 0; i < len(row1); i++ {
		if _, ok := plantNames[row1[i]]; !ok {
			return nil, errors.New("invalid plant code in row 1")
		}
		if _, ok := plantNames[row2[i]]; !ok {
			return nil, errors.New("invalid plant code in row 2")
		}
	}

	// Check for duplicate names
	childrenMap := make(map[string]bool)
	for _, child := range children {
		if childrenMap[child] {
			return nil, errors.New("duplicate child name")
		}
		childrenMap[child] = true
	}

	// Create a new sorted copy of children
	sortedChildren := make([]string, len(children))
	copy(sortedChildren, children)
	sort.Strings(sortedChildren)

	// Assign plants to children
	plantsByChild := make(map[string][]string)
	for i, child := range sortedChildren {
		if i*2 >= len(row1) {
			break
		}

		// Each child gets 2 plants from each row (4 total)
		plants := []string{
			plantNames[row1[i*2]],
			plantNames[row1[i*2+1]],
			plantNames[row2[i*2]],
			plantNames[row2[i*2+1]],
		}
		plantsByChild[child] = plants
	}

	return &Garden{
		plantsByChild: plantsByChild,
	}, nil
}

// Plants returns the plants assigned to a specific child
func (g *Garden) Plants(child string) ([]string, bool) {
	plants, ok := g.plantsByChild[child]
	if !ok {
		return nil, false
	}
	return plants, true
}
