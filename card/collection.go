package card

import (
	"fmt"
	"strconv"
)

type Collection struct {
	Cards []Card
}

// returned collection mode: 36,52 | +2,+36,+100 -> created pack cards length = mode
func NewCollection(mode string) *Collection {
	collect := Collection{}
	switch mode {
	case "36":
		return collect.GenerateCollect36()
	case "52":
		return collect.GenerateCollect52()
	}
	if string(mode[0]) == "+" {
		i, err := strconv.Atoi(mode[1:])
		if err != nil {
			panic("Error mode in constructor Collection")
		}
		return collect.GenerateCollectN(i)
	} else {
		panic("Error params construct!")
	}
}

// add card in collection
func (collect *Collection) Add(card Card) {
	collect.Cards = append(collect.Cards, card)
}

// returned - pop card from collection
func (collect *Collection) Pop() Card {
	ind_end := len(collect.Cards) - 1
	card := collect.Cards[ind_end]
	collect.Cards = collect.Cards[0:ind_end]
	return card
}

// returned count cards in collection
func (collect *Collection) GetCountCards() int {
	return len(collect.Cards)
}

// returned card from collection
func (collect *Collection) GetCard(index int) Card {
	return collect.Cards[index]
}

// returned random card from collection
func (collect *Collection) GetCardRandom() Card {
	index := R(len(collect.Cards))
	card := collect.Cards[index]
	return card
}

// returned card from collection and delete her
func (collect *Collection) GetCardPop(index int) Card {
	card := collect.Cards[index]
	newCollect := collect.Cards[0:index]
	newCollect = append(newCollect, collect.Cards[index:]...)
	collect.Cards = newCollect
	return card
}

// returned random card from collection and delete her
func (collect *Collection) GetCardRandomPop(mode int) Card {
	index := R(len(collect.Cards))
	card := collect.Cards[index]
	newCollect := collect.Cards[0:index]
	newCollect = append(newCollect, collect.Cards[index:]...)
	collect.Cards = newCollect
	return card
}

// shuffle random collection
func (collect *Collection) Suffle() {
	l := len(collect.Cards)
	for i := 0; i < l; i++ {
		random_index := R(l)
		random_index_2 := R(l)
		random_card := collect.Cards[random_index]
		collect.Cards[random_index] = collect.Cards[random_index_2]
		collect.Cards[random_index_2] = random_card
	}
}

// returned collection unique cards - count = 52
func (collect *Collection) GenerateCollect52() *Collection {
	// syms := []int{1,2,3,4}
	syms := []string{"♥", "♦", "♠", "♣"}
	prices := []int{6, 7, 8, 9, 10, 2, 3, 4, 11}
	vals := []string{"2", "3", "4", "5", "6", "7", "8", "9", "0", "V", "D", "K", "T"}
	cards := make([]Card, 0)
	for ind, val := range vals {
		for _, sym := range syms {
			cards = append(cards, *NewCard(string(val), string(sym), prices[ind]))
		}
	}
	return &Collection{cards}
}

// returned collection unique cards - count = 36
func (collect *Collection) GenerateCollect36() *Collection {
	// syms := []int{1,2,3,4}
	syms := []string{"♥", "♦", "♠", "♣"}
	prices := []int{6, 7, 8, 9, 10, 2, 3, 4, 11}
	vals := []string{"6", "7", "8", "9", "0", "V", "D", "K", "T"}
	cards := make([]Card, 0)
	for ind, val := range vals {
		for _, sym := range syms {
			cards = append(cards, *NewCard(string(val), string(sym), prices[ind]))
		}
	}
	return &Collection{cards}
}

// returned collection NOT unique cards - count = N
func (collect *Collection) GenerateCollectN(n int) *Collection {
	// syms := []int{1,2,3,4}
	syms := []string{"♥", "♦", "♠", "♣"}
	vals := []string{"2", "3", "4", "5", "6", "7", "8", "9", "0", "V", "D", "K", "T"}
	prices := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 2, 3, 4, 11}

	cards := make([]Card, 0)
	for i := 0; i < n; i++ {
		rand_s := string(syms[R(len(syms))])
		rand_v := string(vals[R(len(vals))])
		rand_p := prices[R(len(vals))]
		cards = append(cards, *NewCard(string(rand_v), string(rand_s), rand_p))
	}
	return &Collection{cards}
}

// view runes vertical
func (collect *Collection) ViewRunesV() {
	for _, card := range collect.Cards {
		fmt.Println(card.Value, card.Sym)
	}
}

// view runes horizontal
func (collect *Collection) ViewRunesH() {
	line := ""
	spliter := "|"
	for _, card := range collect.Cards {
		line += card.Value + card.Sym + spliter
	}
	fmt.Println(line)
}

// view cards vertical
func (collect *Collection) ViewCardsV() {
	for _, card := range collect.Cards {
		for _, line := range card.Schema {
			fmt.Println(line)
		}
	}
}

// view cards horizontal
func (collect *Collection) ViewCardsH() {
	line1 := ""
	line2 := ""
	line3 := ""
	line4 := ""
	spliter := " "
	for _, card := range collect.Cards {
		line1 += string(card.Schema[0]) + spliter
		line2 += string(card.Schema[1]) + spliter
		line3 += string(card.Schema[2]) + spliter
		line4 += string(card.Schema[3]) + spliter
	}
	fmt.Println(line1)
	fmt.Println(line2)
	fmt.Println(line3)
	fmt.Println(line4)
}

//mode: cardv,cardh,valv,valh,pricev,priceh,symv,symh
func (collect *Collection) View(mode, spliter string, cards ...*Card) {
	line1 := ""
	line2 := ""
	line3 := ""
	line4 := ""
	switch mode {
	case "cardv":
		for _, card := range cards {
			for _, v := range card.Schema {
				fmt.Println(v)
			}
		}
	case "cardh":
		for _, card := range cards {
			line1 += string(card.Schema[0]) + spliter
			line2 += string(card.Schema[1]) + spliter
			line3 += string(card.Schema[2]) + spliter
			line4 += string(card.Schema[3]) + spliter
		}
		fmt.Println(line1)
		fmt.Println(line2)
		fmt.Println(line3)
		fmt.Println(line4)
	case "symv":
		for _, card := range cards {
			fmt.Println(card.Sym)
		}
	case "symh":
		for _, card := range cards {
			fmt.Print(card.Sym, spliter)
		}
	case "pricev":
		for _, card := range cards {
			fmt.Println(card.Price)
		}
	case "priceh":
		for _, card := range cards {
			fmt.Print(card.Price, spliter)
		}
	case "valv":
		for _, card := range cards {
			fmt.Println(card.Value)
		}
	case "valh":
		for _, card := range cards {
			fmt.Print(card.Value + spliter)
		}
	default:
		fmt.Println("Error mode:", mode)
	}
}
