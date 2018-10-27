package dominionizer

// SetName is a type representing name of kingdoms
type SetName int

// Constant enumeration of set names
const (
	Dominion SetName = iota
	Intrigue
	Seaside
	Alchemy
	Prosperity
	Cornucopia
	Hinterlands
	DarkAges
	Guilds
	Adventures
	Empires
	Nocturne
	Renaissance
)

// SetNames is a slice of SetName type for iteration/validation
var SetNames = []SetName{
	Dominion,
	Intrigue,
	Seaside,
	Alchemy,
	Prosperity,
	Cornucopia,
	Hinterlands,
	DarkAges,
	Guilds,
	Adventures,
	Empires,
	Nocturne,
	Renaissance,
}

func (k SetName) String() string {
	names := []string{
		"Dominion",
		"Intrigue",
		"Seaside",
		"Alchemy",
		"Prosperity",
		"Cornucopia",
		"Hinterlands",
		"Dark Ages",
		"Guilds",
		"Adventures",
		"Empires",
		"Nocturne",
		"Renaissance",
	}

	if k < Dominion || k > Renaissance {
		return "Unknown"
	}

	return names[k]
}
