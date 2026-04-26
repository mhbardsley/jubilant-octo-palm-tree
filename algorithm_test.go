package algorithm

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockIndividual struct {
	ID      string
	fitness float64
}

func (m *mockIndividual) Fitness() float64 { return m.fitness }
func (m *mockIndividual) Mutate()          {}

type mockRng struct {
	results []int
	index   int
}

func (m *mockRng) Intn(n int) int {
	result := m.results[m.index]
	m.index++
	return result % n
}

func TestFittest(t *testing.T) {
	t.Run("returns the individual with the highest fitness", func(t *testing.T) {
		population := []*mockIndividual{
			{ID: "1", fitness: 5.0},
			{ID: "2", fitness: 7.0},
			{ID: "3", fitness: 6.0},
		}

		assert.Equal(t, "2", fittest(population).ID)
	})

	t.Run("returns the only individual when the population has one", func(t *testing.T) {
		population := []*mockIndividual{{ID: "1", fitness: 5.0}}

		assert.Equal(t, "1", fittest(population).ID)
	})
}

func TestTournamentSelectRng(t *testing.T) {
	population := []*mockIndividual{
		{ID: "A", fitness: 10.0},
		{ID: "B", fitness: 20.0},
		{ID: "C", fitness: 30.0},
		{ID: "D", fitness: 40.0},
	}

	t.Run("returns the fitter of two distinct picks", func(t *testing.T) {
		selected := tournamentSelectRng(population, 2, &mockRng{results: []int{1, 3}})

		assert.Equal(t, "D", selected.ID)
	})

	t.Run("returns the same individual when both picks collide", func(t *testing.T) {
		selected := tournamentSelectRng(population, 2, &mockRng{results: []int{2, 2}})

		assert.Equal(t, "C", selected.ID)
	})

	t.Run("returns zero value for an empty population", func(t *testing.T) {
		var empty []*mockIndividual

		assert.Nil(t, tournamentSelectRng(empty, 2, &mockRng{}))
	})

	t.Run("k=3 returns the fittest of three picks", func(t *testing.T) {
		// picks A (10), C (30), B (20) → C wins
		selected := tournamentSelectRng(population, 3, &mockRng{results: []int{0, 2, 1}})

		assert.Equal(t, "C", selected.ID)
	})

	t.Run("k<1 is treated as 1", func(t *testing.T) {
		selected := tournamentSelectRng(population, 0, &mockRng{results: []int{1}})

		assert.Equal(t, "B", selected.ID)
	})
}

func TestTopK(t *testing.T) {
	population := []*mockIndividual{
		{ID: "low", fitness: 1},
		{ID: "top1", fitness: 10},
		{ID: "mid", fitness: 5},
		{ID: "top2", fitness: 9},
	}

	elite := topK(population, 2)

	assert.Equal(t, "top1", elite[0].ID)
	assert.Equal(t, "top2", elite[1].ID)
}

func TestNextGenerationElitism(t *testing.T) {
	population := []*mockIndividual{
		{ID: "low", fitness: 1},
		{ID: "mid", fitness: 5},
		{ID: "top1", fitness: 10},
		{ID: "top2", fitness: 9},
	}
	cfg := Config[*mockIndividual]{
		Elitism: 2,
		Crossover: func(a, b *mockIndividual) *mockIndividual {
			return &mockIndividual{ID: "child"}
		},
	}

	next := nextGeneration(population, cfg)

	assert.Len(t, next, 4)
	assert.Equal(t, "top1", next[0].ID)
	assert.Equal(t, "top2", next[1].ID)
	assert.Equal(t, "child", next[2].ID)
	assert.Equal(t, "child", next[3].ID)
}

func TestNextGenerationLocalSearch(t *testing.T) {
	population := []*mockIndividual{
		{ID: "1", fitness: 1},
		{ID: "2", fitness: 2},
		{ID: "3", fitness: 3},
		{ID: "4", fitness: 4},
	}

	t.Run("LocalSearch is invoked on every fresh child", func(t *testing.T) {
		var (
			mu    sync.Mutex
			count int
		)
		cfg := Config[*mockIndividual]{
			Crossover: func(a, b *mockIndividual) *mockIndividual {
				return &mockIndividual{ID: "child"}
			},
			LocalSearch: func(*mockIndividual) {
				mu.Lock()
				count++
				mu.Unlock()
			},
		}

		nextGeneration(population, cfg)

		assert.Equal(t, len(population), count)
	})

	t.Run("LocalSearch skips elite individuals", func(t *testing.T) {
		var (
			mu    sync.Mutex
			count int
		)
		cfg := Config[*mockIndividual]{
			Elitism: 2,
			Crossover: func(a, b *mockIndividual) *mockIndividual {
				return &mockIndividual{ID: "child"}
			},
			LocalSearch: func(*mockIndividual) {
				mu.Lock()
				count++
				mu.Unlock()
			},
		}

		nextGeneration(population, cfg)

		assert.Equal(t, len(population)-cfg.Elitism, count)
	})
}

func TestNextGenerationEliteLocalSearch(t *testing.T) {
	t.Run("EliteLocalSearch is invoked exactly once per generation, on the fittest", func(t *testing.T) {
		population := []*mockIndividual{
			{ID: "1", fitness: 1},
			{ID: "2", fitness: 2},
			{ID: "3", fitness: 3},
			{ID: "4", fitness: 4},
		}
		var seen []string
		cfg := Config[*mockIndividual]{
			Elitism: 1,
			// Children inherit a graded fitness so the elite-pick logic has work to do.
			Crossover: func(a, b *mockIndividual) *mockIndividual {
				return &mockIndividual{ID: "child", fitness: 50}
			},
			EliteLocalSearch: func(m *mockIndividual) {
				seen = append(seen, m.ID)
			},
		}
		next := nextGeneration(population, cfg)

		assert.Len(t, seen, 1, "EliteLocalSearch should run once per generation")
		// Children all have fitness 50; the elite (top1, fitness=4) is at next[0].
		// Children are at next[1..3]. The fittest in next is a child with fitness 50.
		bestFit := next[0].Fitness()
		for _, ind := range next[1:] {
			if ind.Fitness() > bestFit {
				bestFit = ind.Fitness()
			}
		}
		assert.Equal(t, 50.0, bestFit, "test expectation: children were the fittest in next")
		assert.Equal(t, "child", seen[0], "EliteLocalSearch should run on the fittest of next")
	})

	t.Run("EliteLocalSearch composes with LocalSearch: per-child + one for the elite", func(t *testing.T) {
		population := []*mockIndividual{
			{ID: "1", fitness: 1}, {ID: "2", fitness: 2},
			{ID: "3", fitness: 3}, {ID: "4", fitness: 4},
		}
		var (
			mu                    sync.Mutex
			perChild, eliteCalled int
		)
		cfg := Config[*mockIndividual]{
			Elitism: 1,
			Crossover: func(a, b *mockIndividual) *mockIndividual {
				return &mockIndividual{ID: "child", fitness: 99}
			},
			LocalSearch: func(*mockIndividual) {
				mu.Lock()
				perChild++
				mu.Unlock()
			},
			EliteLocalSearch: func(*mockIndividual) {
				eliteCalled++
			},
		}
		nextGeneration(population, cfg)

		assert.Equal(t, len(population)-cfg.Elitism, perChild)
		assert.Equal(t, 1, eliteCalled)
	})
}
