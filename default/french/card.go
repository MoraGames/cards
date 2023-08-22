package french

type FrenchCard struct {
	suit      string
	value     string
	isCovered bool
}

func (c *FrenchCard) Flip() {
	c.isCovered = !c.isCovered
}

func (c *FrenchCard) IsCovered() bool {
	return c.isCovered
}

func (c *FrenchCard) String() string {
	if c.isCovered {
		return "▒▒▒"
	}
	return c.value + " of " + c.suit
}
