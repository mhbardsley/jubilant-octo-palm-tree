package algorithm

import (
	"math/rand"
	"sync"
)

// RunGeneticAlgorithm evolves a population until cfg.ContinuingCondition returns false,
// and returns the fittest individual seen across all generations.
func RunGeneticAlgorithm[T Individual](cfg Config[T]) T {
	population := make([]T, cfg.PopulationSize)
	for i := range population {
		population[i] = cfg.GenerateIndividual()
	}
	best := fittest(population)

	for cfg.ContinuingCondition() {
		population = nextGeneration(population, cfg.Crossover)
		if candidate := fittest(population); candidate.Fitness() > best.Fitness() {
			best = candidate
		}
	}
	return best
}

func nextGeneration[T Individual](population []T, crossover func(T, T) T) []T {
	next := make([]T, len(population))
	var wg sync.WaitGroup
	wg.Add(len(population))
	for i := range population {
		go func(i int) {
			defer wg.Done()
			child := crossover(tournamentSelect(population), tournamentSelect(population))
			child.Mutate()
			next[i] = child
		}(i)
	}
	wg.Wait()
	return next
}

func fittest[T Individual](population []T) T {
	best := population[0]
	for _, candidate := range population[1:] {
		if candidate.Fitness() > best.Fitness() {
			best = candidate
		}
	}
	return best
}

// tournamentSelect picks two random individuals and returns the fitter of the two.
func tournamentSelect[T Individual](population []T) T {
	return tournamentSelectRng(population, defaultRng{})
}

func tournamentSelectRng[T Individual](population []T, r rng) T {
	if len(population) == 0 {
		var zero T
		return zero
	}
	a := population[r.Intn(len(population))]
	b := population[r.Intn(len(population))]
	if b.Fitness() > a.Fitness() {
		return b
	}
	return a
}

type rng interface {
	Intn(int) int
}

type defaultRng struct{}

func (defaultRng) Intn(n int) int { return rand.Intn(n) }
