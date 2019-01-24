package dominionizer

// CreateRandomizer creates an instance of a Randomizer
func CreateRandomizer(id string, sr StateReader, cr CardDataReader) Randomizer {
	r := Randomizer{
		id: id,
		sr: sr,
		cr: cr,
	}

	return r
}
