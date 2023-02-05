package main

import (
	. "Card/card"
	"fmt"
)

func main() {
	f := fmt.Println
	card := NewCardMode("K", 2, 2)
	card.ViewCard()
	card.ViewRune()

	card2 := NewCardMode("D", 3, 3)
	card2.ViewCard()
	card2.ViewRune()

	collect := NewCollection("+7")
	collect.ViewRunesH()
	collect.ViewRunesV()
	collect.ViewCardsH()
	collect.ViewCardsV()

	f(card.IsBigValue(card2))
	f(card.IsSmallValue(card2))
	f(card.IsEquilValue(card2))

	collect.View("cardh", " ", card, card2)
	collect.View("cardh", " ", card, card, card, card2, card2, card2)
	collect.View("cardv", " ", card, card2)
	collect.View("cardv", " ", card, card, card, card2, card2, card2)
	collect.View("valh", " ", card, card2)
	collect.View("valh", " ", card, card, card, card2, card2, card2)
	collect.View("valv", " ", card, card2)
	collect.View("valv", " ", card, card, card, card2, card2, card2)
	collect.View("priceh", " ", card, card2)
	collect.View("priceh", " ", card, card, card, card2, card2, card2)
	collect.View("pricev", " ", card, card2)
	collect.View("pricev", " ", card, card, card, card2, card2, card2)
	collect.View("symh", " ", card, card2)
	collect.View("symh", " ", card, card, card, card2, card2, card2)
	collect.View("symv", " ", card, card2)
	collect.View("symv", " ", card, card, card, card2, card2, card2)

}
