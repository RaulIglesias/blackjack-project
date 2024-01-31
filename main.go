package main

import (
	"fmt"
)

func AnaliseCard(cardName string) int {
	cardValue := map[string]int{
		"ace":   11,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"ten":   10,
		"jack":  10,
		"queen": 10,
		"king":  10,
		"other": 0,
	}
	return cardValue[cardName]
}

// Set a logic of the firts turn
func FirstTurn(card1, card2, dealerCard string) string {

	totalValue := AnaliseCard(card1) + AnaliseCard(card2)
	dealerValue := AnaliseCard(dealerCard)

	switch {
	// Verify pair of ace, Split
	case card1 == "ace" && card2 == "ace":
		return "P"
	//Verify Blackjack, Win or split
	case totalValue == 21 && dealerValue < 10:
		return "W"
	case totalValue == 21:
		return "S"
	//Verify value in (17, 20), Stand
	case totalValue >= 17 && totalValue <= 20:
		return "S"
	//Verify value in (12, 16) and dealerValue, Hit or Stand
	case totalValue >= 12 && totalValue <= 16:
		if dealerValue >= 7 {
			return "H"
		} else {
			return "S"
		}
	//Verify value sum up to 11 or lower, hit
	case totalValue <= 11:
		return "H"

	}

	return "S" // Default: Stand
}

func main() {

	// Cases
	fmt.Println("Case 1: ace, ace, jack")
	fmt.Println("Return:", FirstTurn("ace", "ace", "jack")) // "P"

	fmt.Println("Case 2: ace, king, ace")
	fmt.Println("Return:", FirstTurn("ace", "king", "ace")) // "S"

	fmt.Println("Case 3: five, queen, ace")
	fmt.Println("Return:", FirstTurn("five", "queen", "ace")) // "H"
}
