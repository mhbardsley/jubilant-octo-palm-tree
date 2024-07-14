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

func TestGeneratePopulation(_ *testing.T) {
	// TODO: implement
}

func TestFittestIndividual(t *testing.T) {
	t.Run("fittest individual is first in the slice", func(t *testing.T) {
		individuals := []Individual{
			&mockIndividual{ID: "A", fitness: 10},
			&mockIndividual{ID: "B", fitness: 5},
			&mockIndividual{ID: "C", fitness: 1},
		}

		fittest := fittestIndividual(individuals)
		assert.Equal(t, individuals[0], fittest)
	})

	t.Run("fittest individual is last in the slice", func(t *testing.T) {
		individuals := []Individual{
			&mockIndividual{ID: "A", fitness: 1},
			&mockIndividual{ID: "B", fitness: 5},
			&mockIndividual{ID: "C", fitness: 10},
		}

		fittest := fittestIndividual(individuals)
		assert.Equal(t, individuals[2], fittest)
	})

	t.Run("fittest individual is somewhere else in the slice", func(t *testing.T) {
		individuals := []Individual{
			&mockIndividual{ID: "A", fitness: 1},
			&mockIndividual{ID: "B", fitness: 10},
			&mockIndividual{ID: "C", fitness: 5},
		}

		fittest := fittestIndividual(individuals)
		assert.Equal(t, individuals[1], fittest)
	})

	t.Run("there are two fittest individuals and the first is chosen", func(t *testing.T) {
		individuals := []Individual{
			&mockIndividual{ID: "A", fitness: 10},
			&mockIndividual{ID: "B", fitness: 10},
			&mockIndividual{ID: "C", fitness: 5},
		}

		fittest := fittestIndividual(individuals)
		assert.Equal(t, individuals[0], fittest)
	})

	t.Run("singleton slice returns the expected individual", func(t *testing.T) {
		individuals := []Individual{
			&mockIndividual{ID: "A", fitness: 10},
		}

		fittest := fittestIndividual(individuals)
		assert.Equal(t, individuals[0], fittest)
	})
}

func TestSelectForCrossover(t *testing.T) {
	// TODO: test
}
