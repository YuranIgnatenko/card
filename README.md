> # Card

```
8♠|7♦|D♥|T♣|T♦|2♥|4♥|8♠|7♦|D♥|T♣|T♦|10♥|4♥|
 ___   ___   ___   ___   ___   ___   ___  
|8  | |7  | |D  | |T  | |T  | |9  | |4  | 
| ♠ | | ♦ | | ♥ | | ♣ | | ♦ | | ♥ | | ♥ | 
|__8| |__7| |__D| |__T| |__T| |__9| |__4|
♦ ♠ ♥ ♣ ♦ ♠ ♥ ♣ ♦ ♠ ♥ ♣ ♦ ♠ ♥ ♣ ♦ ♠ ♥ ♣ ♦ 
```

***
Module for simulate Card

> Install

```
go get -u github.com/YuranIgnatenko/Card
```

***

> Main types

``` go
// one card
type Card struct {...}

// pack cards
type Collection struct {
	Cards []Card
}
```

***

> Easy use

``` go
package main

import (
	. "github.com/YuranIgnatenko/Card/card"
	"fmt"
)

func main() {
    // "K"-value in card, 2-type suit, 2-price card
    // Create card
    card := NewCardMode("K", 2, 2)

    // View Card
    card.ViewCard()

    // View Rune
    card.ViewRune()
}

```
> Out terminal
```
<!-- Card -->
 ___ 
|K  |
| ♦ |
|__K|

<!-- Rune -->
K♦
```

***
> Select suits card
``` go
// use number:
// 1:"♥", 2:"♦", 3:"♠", 4:"♣"
card1 := NewCardMode("D", 3, 3)
card2 := NewCardMode("10", 3, 10)

// or

// use symbol:
// "♥", "♦", "♠", "♣"
card1 := NewCard("D", "♣", 3)
card2 := NewCard("10","♥", 10)

// All return type [*Card]
```
***
> Compare cards (>,<,=)
``` go
card_k := NewCardMode("K", 3, 13) 
card_10 := NewCardMode("10", 3, 10)

card_10.IsBigValue(card_k) // card_10 > card_k = false
card_10.IsSmallValue(card_k) // card_10 < card_k = true
card_10.IsEquilValue(card_k) // card_10 == card_k = false

// or

card_k.IsBigValue(card_10) // card_10 > card_k = true
card_k.IsSmallValue(card_10) // card_10 < card_k = false
card_k.IsEquilValue(card_10) // card_10 == card_k = false

```
***
> Pack cards
``` go
// generate rigth pack cards (no dublicate)
collect := NewCollection("36")
collect := NewCollection("52")

// generate your count cards (possible dublicate!)
collect := NewCollection("+7")
collect := NewCollection("+143")
collect := NewCollection("+9999")

// possible alternative generate pack-cards
collect := GenerateCollect36()
collect := GenerateCollect52()
collect := GenerateCollectN(count int)

// All returns type [*Collection]
```
***
> Gameplay for pack cards

``` go
card_k := NewCardMode("K", 3, 13) 
collect := NewCollection("36")

// Get len pack-cards (count cards)
collect.GetCountCards() // return 35

// Add card in pack-cards
collect.Add(card_k)

// Pop last card from pack-cards
collect.Pop() // return *Card("K",3,13)

// Get copy random-card from pack-cards
collect.GetCardRandom() // return card-copy

// Get copy card from pack-cards
collect.GetCard(index int) // return card-copy

// Get original card from pack-cards
collect.GetCardPop(index int) // return original card

// Get original random-card from pack-cards
collect.GetCardRandomPop() // return original card 

// Suffle pack-cards
collect.Suffle()
```
****

> View pack-cards
``` go
collect := NewCollection("+2")

collect.ViewRunesH()
8♠|7♦|


collect.ViewRunesV()
8 ♠
7 ♦


collect.ViewCardsH()
 ___   ___  
|8  | |7  |
| ♠ | | ♦ | 
|__8| |__7|


collect.ViewCardsV()
 ___ 
|8  |
| ♠ |
|__8|
 ___ 
|7  |
| ♦ |
|__7|


// Developing method
// mode : cardv,cardh,valv,valh,pricev,priceh,symv,symh
collect.View(mode, spliter, ...cards)
collect.View("cardh", " ", card, card2)
```
***

