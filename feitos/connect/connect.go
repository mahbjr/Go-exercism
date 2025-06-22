package connect

import (
	"strings"
)

// ResultOf determines the winner of a Connect board
func ResultOf(board []string) (string, error) {
	// Clean up the board representation - remove spaces and indentation
	cleanBoard := make([]string, len(board))
	for i, row := range board {
		cleanBoard[i] = strings.ReplaceAll(strings.TrimSpace(row), " ", "")
	}

	// Check for X win (left to right)
	if hasPath(cleanBoard, 'X', true) {
		return "X", nil
	}

	// Check for O win (top to bottom)
	if hasPath(cleanBoard, 'O', false) {
		return "O", nil
	}

	// No winner
	return "", nil
}

// Check if player has a winning path
func hasPath(board []string, player byte, isX bool) bool {
	if len(board) == 0 {
		return false
	}

	// Special case for 1x1 board
	if len(board) == 1 && len(board[0]) == 1 {
		return board[0][0] == player
	}

	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[i]))
	}

	// For X: start from left edge, check if we can reach right edge
	if isX {
		for i := 0; i < len(board); i++ {
			if len(board[i]) > 0 && board[i][0] == player {
				if dfs(board, visited, i, 0, player, isX) {
					return true
				}
			}
		}
	} else {
		// For O: start from top edge, check if we can reach bottom edge
		for j := 0; j < len(board[0]); j++ {
			if board[0][j] == player {
				if dfs(board, visited, 0, j, player, isX) {
					return true
				}
			}
		}
	}

	return false
}

// DFS to find a path
func dfs(board []string, visited [][]bool, i, j int, player byte, isX bool) bool {
	// Check if position is out of bounds
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) {
		return false
	}

	// Check if position has been visited or isn't player's stone
	if visited[i][j] || board[i][j] != player {
		return false
	}

	// Mark as visited
	visited[i][j] = true

	// Check if we've reached the winning edge
	if (isX && j == len(board[i])-1) || (!isX && i == len(board)-1) {
		return true
	}

	// Hexagonal neighbors - six possible directions
	directions := [][2]int{
		{-1, 0},  // up
		{-1, 1},  // up-right
		{0, -1},  // left
		{0, 1},   // right
		{1, -1},  // down-left
		{1, 0},   // down
	}

	// Explore all neighbors
	for _, dir := range directions {
		ni, nj := i+dir[0], j+dir[1]
		if dfs(board, visited, ni, nj, player, isX) {
			return true
		}
	}

	return false
}