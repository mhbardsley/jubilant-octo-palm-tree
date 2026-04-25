package algorithm

// Individual is a candidate solution that can be ranked by fitness and mutated in place.
type Individual interface {
	Fitness() float64
	Mutate()
}

// Config describes a genetic algorithm: how to build individuals, how to combine
// them, and when to stop.
//
// Conventions for the user-supplied operators:
//   - Mutate modifies the receiver in place.
//   - Crossover returns a fresh T and must not modify or alias either parent.
//   - LocalSearch, if set, modifies its argument in place (like Mutate).
type Config[T Individual] struct {
	PopulationSize      int
	GenerateIndividual  func() T
	Crossover           func(T, T) T
	ContinuingCondition func() bool

	// Elitism copies the top K individuals verbatim into each new generation,
	// protecting the best genomes from being lost to crossover and mutation.
	// Zero (the default) disables elitism and reproduces the original behavior.
	Elitism int

	// TournamentSize controls selection pressure: each tournament picks
	// TournamentSize individuals at random and returns the fittest.
	// Zero (the default) is treated as 2.
	TournamentSize int

	// LocalSearch, if non-nil, is called on each freshly bred child after
	// crossover and Mutate, turning the algorithm into a memetic / Lamarckian
	// GA. It must improve its argument in place. Elite individuals are not
	// passed through LocalSearch (they were already polished in the generation
	// that produced them).
	LocalSearch func(T)
}
