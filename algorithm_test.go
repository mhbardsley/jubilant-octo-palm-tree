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
