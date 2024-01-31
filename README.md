# Go Blackjack Simple Logic

This Go program implements a simple logic for the first round in a Blackjack game. It determines the optimal play (Hit, Stand, Win, or Split) based on the player's cards and the dealer's revealed card.

## Key Concepts

### Function: `AnaliseCard`

The `AnaliseCard` function takes the name of a card as a parameter and returns its value based on a predefined map. It uses a map to associate card names with their respective values.

### Function: `FirstTurn`

The `FirstTurn` function takes three parameters: `card1`, `card2`, and `dealerCard`. It calculates the total value of the player's two initial cards (`card1` and `card2`) and the value of the dealer's revealed card (`dealerCard`). It then uses a series of conditions to determine the best play for the player:

- **Pair of Aces (`"P"`):** If the player's initial cards form a pair of Aces, the function recommends splitting.

- **Blackjack (`"W"` or `"S"`):** If the total value is 21 and the dealer's revealed card is less than 10, it recommends winning (`"W"`). Otherwise, it recommends standing (`"S"`).

- **Stand (`"S"`):** If the total value is between 17 and 20, it recommends standing.

- **Hit or Stand (`"H"` or `"S"`):** If the total value is between 12 and 16 and the dealer's revealed card is 7 or more, it recommends hitting (`"H"`). Otherwise, it recommends standing (`"S"`).

- **Hit (`"H"`):** If the total value is 11 or less, it recommends hitting (`"H"`).

### Function: `main`

The `main` function provides test cases to demonstrate the `FirstTurn` logic in different scenarios.

## How to Run

To run the program, use the following command in your terminal:

```bash
go run main.go
``` 

## Contributions

This program is a simple implementation for educational purposes. Contributions and suggestions for improvement are welcome.


## Example Output 

```
Case 1: ace, ace, jack
Return: "P"

Case 2: ace, king, ace
Return: "S"

Case 3: five, queen, ace
Return: "H"