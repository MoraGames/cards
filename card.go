package cards

type Card interface {
	Flip()
	IsCovered() bool
	String() string
}
