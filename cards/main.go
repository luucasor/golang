package main

import "fmt"

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")
	cards.print()

	var ls laptopSize = 35.4
	fmt.Println(ls.getSizeOfLaptop())
}

func newCard() string {
	return "Five of Diamonds"
}
