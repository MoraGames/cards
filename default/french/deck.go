package french

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/MoraGames/cards"
)

type FrenchDeck struct {
	size  int
	cards []cards.Card
}

func DefaultFrenchDeck() *FrenchDeck {
	var frenchDeck FrenchDeck

	frenchDeck.size = 52
	for i := 0; i < frenchDeck.size; i++ {

		var cardSuit, cardValue string
		switch i / 13 {
		case 0:
			cardSuit = "♥"
		case 1:
			cardSuit = "♦"
		case 2:
			cardSuit = "♣"
		case 3:
			cardSuit = "♠"
		}
		switch i % 13 {
		case 0:
			cardValue = " A"
		case 1, 2, 3, 4, 5, 6, 7, 8:
			cardValue = " " + strconv.Itoa(i%13+1)
		case 9:
			cardValue = "10"
		case 10:
			cardValue = " J"
		case 11:
			cardValue = " Q"
		case 12:
			cardValue = " K"
		}

		frenchDeck.cards = append(frenchDeck.cards, &FrenchCard{
			suit:      cardSuit,
			value:     cardValue,
			isCovered: true,
		})
	}
	return &frenchDeck
}

func (d *FrenchDeck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) { d.cards[i], d.cards[j] = d.cards[j], d.cards[i] })
}

func (d *FrenchDeck) Raise() cards.Card {
	rand.Seed(time.Now().UnixNano())
	splitPoint := rand.Intn(d.size-1) + 1
	splitCard := d.cards[splitPoint-1]

	var newDeck *FrenchDeck
	for i := 0; i < d.size; i++ {
		newDeck.cards[(splitPoint+i)%d.size] = d.cards[i]
	}
	d.cards = newDeck.cards

	return splitCard
}

func (d *FrenchDeck) Deal(n int) ([]cards.Card, error) {
	if n > d.size || n < 0 {
		return nil, fmt.Errorf("Number of cards to deal is out of range")
	}

	var dealedCards []cards.Card
	for i := 0; i < n; i++ {
		dealedCards = append(dealedCards, d.cards[i])
	}

	d.cards = d.cards[n:]
	d.size -= n
	return dealedCards, nil
}

func (d *FrenchDeck) Insert(c cards.Card, n int) error {
	if n > d.size || n < 0 {
		return fmt.Errorf("Insert index is out of range")
	}

	var newDeck *FrenchDeck
	newDeck.cards = append(newDeck.cards, d.cards[:n]...)
	newDeck.cards = append(newDeck.cards, c)
	newDeck.cards = append(newDeck.cards, d.cards[n:]...)
	newDeck.size = d.size + 1
	d = newDeck
	return nil
}

func (d *FrenchDeck) Flip() error {
	for _, c := range d.cards {
		c.Flip()
	}
	return nil
}

func (d *FrenchDeck) IsCovered() bool {
	return d.cards[0].IsCovered()
}

func (d *FrenchDeck) String() string {
	return d.cards[0].String()
}

func (d *FrenchDeck) Size() int {
	return d.size
}
