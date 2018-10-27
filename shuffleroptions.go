package dominionizer

// ShufflerOptions is a type which contains options to limit the generation of random kingdoms
type ShufflerOptions struct {
	IncludedSets     map[SetName]bool
	BlacklistedCards map[string]bool
	MustIncludeCards []string
	TypeRules        map[CardType]int
}
