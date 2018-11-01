package dominionizer

// SetName is a type representing name of kingdoms
type SetName int

// Constant enumeration of set names
const (
	SetDominion SetName = iota
	SetIntrigue
	SetSeaside
	SetAlchemy
	SetProsperity
	SetCornucopia
	SetHinterlands
	SetDarkAges
	SetGuilds
	SetAdventures
	SetEmpires
	SetNocturne
	SetRenaissance
)

// SetNames is a slice of SetName type for iteration/validation
var SetNames = []SetName{
	SetDominion,
	SetIntrigue,
	SetSeaside,
	SetAlchemy,
	SetProsperity,
	SetCornucopia,
	SetHinterlands,
	SetDarkAges,
	SetGuilds,
	SetAdventures,
	SetEmpires,
	SetNocturne,
	SetRenaissance,
}

var stringNames = []string{
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

func (k SetName) String() string {
	if k < SetDominion || k > SetRenaissance {
		return "Unknown"
	}

	return stringNames[k]
}

// SetNameFromString takes a string value and converts it to appropriate SetName constant
func SetNameFromString(s string) SetName {
	for i, sn := range stringNames {
		if sn == s {
			return SetNames[i]
		}
	}

	return SetDominion
}
