package algorithm

import (
	"math/rand"
	"sync"
	"time"
)

// RunGeneticAlgorithm is the main function that runs the genetic algorithm.
func RunGeneticAlgorithm[T Individual](config Config[T]) T {
	population, bestIndividual := initializePopulation(config.PopulationSize, config.AlgorithmConfig.GenerateIndividual)

	for config.AlgorithmConfig.ContinuingCondition() {
		population = runIteration(population, config.AlgorithmConfig.GenerateCrossover)
		bestIndividual = updateBestIndividual(population, bestIndividual)
	}

	return bestIndividual
}

// initializePopulation creates an initial population and returns it along with the initial best individual.
func initializePopulation[T Individual](populationSize int, generateIndividual func() T) ([]T, T) {
	population := make([]T, populationSize)
	for i := 0; i < populationSize; i++ {
		population[i] = generateIndividual()
	}

	bestIndividual := population[0]
	for _, individual := range population[1:] {
		if individual.Fitness() > bestIndividual.Fitness() {
			bestIndividual = individual
		}
	}

	return population, bestIndividual
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

// updateBestIndividual checks if the new population contains a better individual and updates the best one found so far.
func updateBestIndividual[T Individual](population []T, bestIndividual T) T {
	bestFitness := bestIndividual.Fitness()

	for _, individual := range population {
		if individual.Fitness() > bestFitness {
			bestIndividual = individual
			bestFitness = individual.Fitness()
		}
	}

	return bestIndividual
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
		if contender.Fitness() > best.Fitness() {
			best = contender
		}
	}

	return best
}
