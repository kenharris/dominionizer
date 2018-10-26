package dominionizer

// ShufflerOptions is a type which contains options to limit the generation of random kingdoms
type ShufflerOptions struct {
	IncludedKingdoms map[SetName]bool
	BlacklistedCards []string
	MustIncludeCards []string
}
