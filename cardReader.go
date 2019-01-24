package dominionizer

// CardReader is an interface that allows for reading randomizer cards
type CardReader interface {
	ReadCards() []Card
}
