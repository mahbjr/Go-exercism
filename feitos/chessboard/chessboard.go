package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	// Get the file from the chessboard
	f, exists := cb[file]
	if !exists {
		return 0
	}

	// Count occupied squares in the file
	count := 0
	for _, occupied := range f {
		if occupied {
			count++
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	// Check if rank is valid (1-8)
	if rank < 1 || rank > 8 {
		return 0
	}

	// Convert to 0-based index
	rankIndex := rank - 1

	// Count occupied squares in the rank
	count := 0
	// Loop through all files
	for _, file := range []string{"A", "B", "C", "D", "E", "F", "G", "H"} {
		if f, exists := cb[file]; exists && rankIndex < len(f) {
			if f[rankIndex] {
				count++
			}
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	// A standard chessboard has 64 squares (8x8)
	// But we'll calculate it by counting the squares in each file

	total := 0
	for _, file := range cb {
		total += len(file)
	}
	return total
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	// Count all occupied squares across all files
	count := 0
	for _, file := range cb {
		for _, occupied := range file {
			if occupied {
				count++
			}
		}
	}
	return count
}
