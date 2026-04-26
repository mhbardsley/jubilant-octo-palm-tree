package algorithm

import (
	"math/rand"
	"sort"
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
		population = nextGeneration(population, cfg)
		if candidate := fittest(population); candidate.Fitness() > best.Fitness() {
			best = candidate
		}
	}
	return best
}

func nextGeneration[T Individual](population []T, cfg Config[T]) []T {
	n := len(population)
	next := make([]T, n)

	elite := cfg.Elitism
	if elite < 0 {
		elite = 0
	}
	if elite > n {
		elite = n
	}
	if elite > 0 {
		copy(next[:elite], topK(population, elite))
	}

	k := cfg.TournamentSize
	if k < 1 {
		k = 2
	}

	var wg sync.WaitGroup
	for i := elite; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			child := cfg.Crossover(tournamentSelect(population, k), tournamentSelect(population, k))
			child.Mutate()
			if cfg.LocalSearch != nil {
				cfg.LocalSearch(child)
			}
			next[i] = child
		}(i)
	}
	wg.Wait()

	// One last polish on the new generation's fittest, if requested. Runs
	// sequentially after children are bred so the local search sees the final
	// state of next. Modifying the slice element in place propagates because
	// Individual is typically a pointer type.
	if cfg.EliteLocalSearch != nil {
		bestIdx := 0
		bestFit := next[0].Fitness()
		for i := 1; i < n; i++ {
			if f := next[i].Fitness(); f > bestFit {
				bestIdx = i
				bestFit = f
			}
		}
		cfg.EliteLocalSearch(next[bestIdx])
	}
	return next
}

// topK returns the k fittest individuals, fittest first.
func topK[T Individual](population []T, k int) []T {
	type scored struct {
		ind T
		fit float64
	}
	s := make([]scored, len(population))
	for i, ind := range population {
		s[i] = scored{ind, ind.Fitness()}
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i].fit > s[j].fit
	})
	out := make([]T, k)
	for i := 0; i < k; i++ {
		out[i] = s[i].ind
	}
	return out
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

// tournamentSelect picks k random individuals and returns the fittest.
func tournamentSelect[T Individual](population []T, k int) T {
	return tournamentSelectRng(population, k, defaultRng{})
}

func tournamentSelectRng[T Individual](population []T, k int, r rng) T {
	if len(population) == 0 {
		var zero T
		return zero
	}
	if k < 1 {
		k = 1
	}
	best := population[r.Intn(len(population))]
	for i := 1; i < k; i++ {
		candidate := population[r.Intn(len(population))]
		if candidate.Fitness() > best.Fitness() {
			best = candidate
		}
	}
	return best
}

type rng interface {
	Intn(int) int
}

type defaultRng struct{}

func (defaultRng) Intn(n int) int { return rand.Intn(n) }
