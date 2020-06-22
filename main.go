package main

// return string detected
func newCard() string {
  return "Ace of Spades"
}

func main() {
    // card := newCard() // function assignation

    // Slice
    cards := newDeck()

    // Save to File
    cards.saveToFile("cards.txt")

    // print
    cards.print()

    // deals
    hand, remainingDeck := deal(cards,5)
    hand.print()
    remainingDeck.print()

    // Deck from File
    newDeck := newDeckFromFile("cards.txt")
    newDeck.print()

    // shuffle
    cards.shuffe()
    cards.print()
}
