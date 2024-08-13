package algorithm

type GeneticAlgorithm[T Individual] interface {
	GenerateIndividual() T
	GenerateCrossover(T, T) T
	ContinuingCondition() bool
}

type Individual interface {
	Fitness() float64
	Mutate()
}

type Config[T Individual] struct {
	PopulationSize  int
	AlgorithmConfig GeneticAlgorithm[T]
}

type rng interface {
	Intn(int) int
}
