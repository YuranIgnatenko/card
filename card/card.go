package card

import (
	"fmt"
)

type Card struct {
	Schema         []string
	Value          string
	Sym            string
	Price          int
	possibleVals   []string
	possibleSyms   []string
	possiblePrices []int
}

// returned Card ("8", "♥", 8)
func NewCard(value, sym string, price int) *Card {
	card := Card{}
	card.Schema = []string{
		" ___ ",
		"|" + value + "  |",
		"| " + sym + " |",
		"|__" + value + "|"}
	card.Value = value
	card.Sym = sym
	card.Price = price
	card.possiblePrices = []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 2, 3, 4, 11}
	card.possibleVals = []string{"2", "3", "4", "5", "6", "7", "8", "9", "0", "V", "D", "K", "T"}
	card.possibleSyms = []string{"♥", "♦", "♠", "♣"}
	return &card
}

// returned Card ("8", symMode(1:"♥", 2:"♦", 3:"♠", 4:"♣"), 8)
func NewCardMode(value string, symMode int, price int) *Card {
	if symMode < 0 && symMode > 5 {
		panic("Error mode! Need number of range 1...4")
	}

	selectSym := []string{"♥", "♦", "♠", "♣"}
	// selectVal := []string{"2","3","4","5","6","7","8","9","0","V","D","K","T"}
	sym := selectSym[symMode-1]

	card := Card{}
	card.Schema = []string{
		" ___ ",
		"|" + value + "  |",
		"| " + sym + " |",
		"|__" + value + "|"}
	card.Value = value
	card.Sym = sym
	card.Price = price
	return &card
}

// returned bool check is big value card <> card2
func (card *Card) IsBigValue(card2 *Card) bool {
	indcard := 0
	indcard2 := 0
	for i, val := range card.possibleVals {
		if card.Value == val {
			indcard = i
		}
		if card2.Value == val {
			indcard2 = i
		}
	}
	return indcard > indcard2
}

// returned bool check is small value card <> card2
func (card *Card) IsSmallValue(card2 *Card) bool {
	indcard := 0
	indcard2 := 0
	for i, val := range card.possibleVals {
		if card.Value == val {
			indcard = i
		}
		if card2.Value == val {
			indcard2 = i
		}
	}
	return indcard < indcard2
}

// returned bool check is equil value card <> card2
func (card *Card) IsEquilValue(card2 *Card) bool {
	indcard := 0
	indcard2 := 0
	for i, val := range card.possibleVals {
		if card.Value == val {
			indcard = i
		}
		if card2.Value == val {
			indcard2 = i
		}
	}
	return indcard == indcard2
}

// view rune card
func (card *Card) ViewRune() {
	fmt.Println(card.Value + card.Sym)
}

// view card
func (card *Card) ViewCard() {
	for _, line := range card.Schema {
		fmt.Println(line)
	}
}
