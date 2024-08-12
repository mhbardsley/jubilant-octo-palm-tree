package algorithm

import (
	"math/rand"
	"sync"
	"time"
)

func RunGeneticAlgorithm[T Individual](config Config[T]) T {
	population := make([]T, config.PopulationSize)
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
	var wg sync.WaitGroup
	for i := range len(population) {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ind1 := selectForCrossover(population)
			ind2 := selectForCrossover(population)
			child := crossoverFunc(ind1, ind2)
			child.Mutate()
			newPop[i] = child
		}(i)
	}
	wg.Wait()
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
	if len(population) == 0 {
		var zero T
		return zero
	}

	k := 2
	best := population[rng.Intn(len(population))]

	for i := 1; i < k; i++ {
		contender := population[rng.Intn(len(population))]
		if contender.GetFitness() > best.GetFitness() {
			best = contender
		}
	}

	return best
}
