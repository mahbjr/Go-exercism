package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	// Calculate the score of the player's hand
	cardValue1 := ParseCard(card1)
	cardValue2 := ParseCard(card2)
	playerScore := cardValue1 + cardValue2

	dealerValue := ParseCard(dealerCard)

	// Check for blackjack (ace + 10-value card)
	isBlackjack := playerScore == 21

	// Check for pair of aces
	isPairOfAces := card1 == "ace" && card2 == "ace"

	// Rules based on the test cases:

	// Rule 1: If the player has blackjack and the dealer doesn't have an ace or a 10-value card,
	// the player automatically wins
	if isBlackjack && dealerValue < 10 {
		return "W" // Win
	}

	// Rule 2: If the player has blackjack and the dealer has an ace or a 10-value card,
	// the player stands
	if isBlackjack {
		return "S" // Stand
	}

	// Rule 3: With a pair of aces, the player splits
	if isPairOfAces {
		return "P" // Split
	}

	// Rule 4: With a score of 17 or higher, the player stands
	if playerScore >= 17 {
		return "S" // Stand
	}

	// Rule 5: With a score of 16 or lower and the dealer showing a card with value 6 or lower,
	// the player stands
	if playerScore >= 12 && playerScore <= 16 && dealerValue <= 6 {
		return "S" // Stand
	}

	// Rule 6: With a score of 11 or lower, the player always hits
	// Rule 7: In all other cases, the player hits
	return "H" // Hit
}
