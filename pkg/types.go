package pkg

type GeneticAlgorithm interface {
	GenerateIndividual() Individual
	ContinuingCondition() bool
	Individual
}

type Individual interface {
	GetFitness() float64
	Mutate()
}

type Config struct {
	PopulationSize uint
	AlgorithmConfig GeneticAlgorithm
}