package algorithm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockIndividual struct {
	ID      string
	fitness float64
}

func (m *mockIndividual) GetFitness() float64 {
	return m.fitness
}

func (m *mockIndividual) Mutate() {
	// Mock implementation
}

func (m *mockIndividual) GenerateCrossoverWith(other Individual) Individual {
	// Mock implementation
	return m
}

type mockRng struct {
	results []int
	index   int
}

func (m *mockRng) Intn(n int) int {
	result := m.results[m.index]
	m.index++
	return result % n
}

func TestFittestIndividual(t *testing.T) {
	t.Run("fittest individual is first in the slice", func(t *testing.T) {
		individuals := []*mockIndividual{
			&mockIndividual{ID: "A", fitness: 10},
			&mockIndividual{ID: "B", fitness: 5},
			&mockIndividual{ID: "C", fitness: 1},
		}

		fittest := fittestIndividual(individuals)
		assert.Equal(t, "A", fittest.ID)
	})

	t.Run("fittest individual is last in the slice", func(t *testing.T) {
		individuals := []*mockIndividual{
			&mockIndividual{ID: "A", fitness: 1},
			&mockIndividual{ID: "B", fitness: 5},
			&mockIndividual{ID: "C", fitness: 10},
		}

		fittest := fittestIndividual(individuals)
		assert.Equal(t, "C", fittest.ID)
	})

	t.Run("fittest individual is somewhere else in the slice", func(t *testing.T) {
		individuals := []*mockIndividual{
			&mockIndividual{ID: "A", fitness: 1},
			&mockIndividual{ID: "B", fitness: 10},
			&mockIndividual{ID: "C", fitness: 5},
		}

		fittest := fittestIndividual(individuals)
		assert.Equal(t, "B", fittest.ID)
	})

	t.Run("there are two fittest individuals and the first is chosen", func(t *testing.T) {
		individuals := []*mockIndividual{
			&mockIndividual{ID: "A", fitness: 10},
			&mockIndividual{ID: "B", fitness: 10},
			&mockIndividual{ID: "C", fitness: 5},
		}

		fittest := fittestIndividual(individuals)
		assert.Equal(t, "A", fittest.ID)
	})

	t.Run("singleton slice returns the expected individual", func(t *testing.T) {
		individuals := []*mockIndividual{
			&mockIndividual{ID: "A", fitness: 10},
		}

		fittest := fittestIndividual(individuals)
		assert.Equal(t, "A", fittest.ID)
	})
}

func TestSelectForCrossoverRng(t *testing.T) {
	population := []*mockIndividual{
		{ID: "A", fitness: 10.0},
		{ID: "B", fitness: 20.0},
		{ID: "C", fitness: 30.0},
		{ID: "D", fitness: 40.0},
	}

	t.Run("select the first individual if it is fitter", func(t *testing.T) {
		mock := &mockRng{
			results: []int{1, 0},
		}

		selected := selectForCrossoverRng(population, mock)

		assert.Equal(t, "B", selected.ID)
	})

	t.Run("select the second individual if it is fitter", func(t *testing.T) {
		mock := &mockRng{
			results: []int{1, 3}, // Will select individuals 1 and 3
		}

		selected := selectForCrossoverRng(population, mock)

		assert.Equal(t, "D", selected.ID)
	})

	t.Run("select individual with higher fitness even if indices are repeated", func(t *testing.T) {
		mock := &mockRng{
			results: []int{2, 2},
		}

		selected := selectForCrossoverRng(population, mock)

		assert.Equal(t, "C", selected.ID)
	})

	t.Run("handle nil population", func(t *testing.T) {
		mock := &mockRng{
			results: []int{0},
		}

		var nilPopulation []*mockIndividual
		selected := selectForCrossoverRng(nilPopulation, mock)

		assert.Nil(t, selected)
	})

	t.Run("handle empty population", func(t *testing.T) {
		mock := &mockRng{
			results: []int{0},
		}

		emptyPopulation := []*mockIndividual{}
		selected := selectForCrossoverRng(emptyPopulation, mock)

		assert.Nil(t, selected)
	})
}
