package dominionizer

// CardRepositoryReader is an interface which reads Card structs from a source
type CardDataReader interface {
	ReadCards() []Card
}
