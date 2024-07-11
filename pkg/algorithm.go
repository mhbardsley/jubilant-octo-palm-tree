package pkg

func RunGeneticAlgorithm(config Config) Individual {
	population := make([]Individual, config.PopulationSize)
	for i := range config.PopulationSize {
		population[i] = GenerateIndividual()
	}
	for config.AlgorithmConfig.ContinuingCondition() {
		population = runIteration(population)
	}
	return fittestIndividual(population)
}

func runIteration(population []Individual) []Individual {
	newPop := make([]Individual, len(population))
	for i := range len(population) {
		ind1, ind2 := SelectForCrossover(population)
		child := GenerateCrossover(ind1, ind2)  // TODO: test that this is commutative
		newPop[i] = child
	}
	return newPop
}

// TODO: assertion that new populations have the same size in runIteration