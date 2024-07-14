package algorithm

type GeneticAlgorithm[T Individual] interface {
	GenerateIndividual() T
	GenerateCrossover(T, T) T
	ContinuingCondition() bool
}

type Individual interface {
	GetFitness() float64
	Mutate()
}

type Config struct {
	PopulationSize  int
	AlgorithmConfig GeneticAlgorithm[Individual]
}
