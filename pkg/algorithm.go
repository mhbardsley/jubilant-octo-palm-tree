package pkg

func RunGeneticAlgorithm(config Config) Individual {
	population := make([]Individual, config.PopulationSize)
	for i := range config.PopulationSize {
		population[i] = config.AlgorithmConfig.GenerateIndividual()
	}
	for config.AlgorithmConfig.ContinuingCondition() {
		population = runIteration(population, config.AlgorithmConfig.GenerateCrossover)
	}
	return fittestIndividual(population)
}

func runIteration[T Individual](population []T, crossoverFunc func(T, T) T) []T {
	newPop := make([]T, len(population))
	for i := range len(population) { // TODO: big win for concurrency here
		ind1, ind2 := selectForCrossover(population)
		child := crossoverFunc(ind1, ind2) // TODO: test that this is commutative
		newPop[i] = child
	}
	return newPop
}

func fittestIndividual[T Individual](population []T) T {
	// TODO: can we make an assumption that it will be nonempty?

	fittest := population[0]
	for _, individual := range population[1:] {
		if individual.GetFitness() > fittest.GetFitness() {
			fittest = individual
		}
	}

	return fittest
}

func selectForCrossover[T Individual](population []T) (T, T) {
	return population[0], population[1]
}
