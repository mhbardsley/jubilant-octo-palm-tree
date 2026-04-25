package algorithm

// Individual is a candidate solution that can be ranked by fitness and mutated in place.
type Individual interface {
	Fitness() float64
	Mutate()
}

// Config describes a genetic algorithm: how to build individuals, how to combine them, and when to stop.
type Config[T Individual] struct {
	PopulationSize      int
	GenerateIndividual  func() T
	Crossover           func(T, T) T
	ContinuingCondition func() bool
}
