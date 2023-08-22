package cards

type Deck interface {
	Shuffle()
	Raise() Card
	Deal(int) ([]Card, error)
	Insert(Card, int) error
	Flip() error
	IsCovered() bool
	String() string
	Size() int
}
