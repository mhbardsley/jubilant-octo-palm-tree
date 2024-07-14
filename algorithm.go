package algorithm

import (
	"cmp"
	"math/rand"
	"slices"
	"time"
)

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
	slices.SortFunc(population, func(a, b T) int {
		return cmp.Compare(a.GetFitness(), b.GetFitness())
	})
	newPop := make([]T, len(population))
	for i := range len(population) {
		ind1 := selectForCrossover(population)
		ind2 := selectForCrossover(population)
		child := crossoverFunc(ind1, ind2)
		child.Mutate()
		newPop[i] = child
	}
	return newPop
}

func fittestIndividual[T Individual](population []T) T {
	fittest := population[0]
	for _, individual := range population[1:] {
		if individual.GetFitness() > fittest.GetFitness() {
			fittest = individual
		}
	}

	return fittest
}

func selectForCrossover[T Individual](population []T) T {
	// we need to assume that this is already sorted, ascending, by fitness, as we use a ranked crossover function
	// see if we can have a test for this
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	return selectForCrossoverRng(population, rng)
}

func selectForCrossoverRng[T Individual](population []T, rng rng) T {
	populationSize := len(population)
	// Calculate cumulative probabilities based on 1, 2, 3, ..., n
	cumulativeProb := make([]int, populationSize)
	cumulativeProb[0] = 1 // prob(element 0) = 1
	for i := 1; i < populationSize; i++ {
		cumulativeProb[i] = cumulativeProb[i-1] + (i + 1)
	}

	// Generate a random number between 0 and sum of cumulative probabilities
	totalProb := cumulativeProb[populationSize-1]
	randomNum := rng.Intn(totalProb)

	// Find the element corresponding to the selected cumulative probability
	for i := 0; i < populationSize; i++ {
		if randomNum < cumulativeProb[i] {
			return population[i]
		}
	}

	// Fallback (shouldn't happen with correct random generation and cumulative probabilities)
	return population[populationSize-1]
}
