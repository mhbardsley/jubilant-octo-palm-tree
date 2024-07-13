package pkg  // TODO: could this have a better name?

import (
	"testing"
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
			&MockIndividual{ID: "A", fitness: 10},
			&MockIndividual{ID: "B", fitness: 5},
			&MockIndividual{ID: "C", fitness: 1},
		}

		fittest := fittestIndividual(individuals)
		assert.Equal(t, "A", fittest.ID)
	})

	t.Run("fittest individual is last in the slice", func(t *testing.T) {
		individuals := []Individual{
			&MockIndividual{ID: "A", fitness: 1},
			&MockIndividual{ID: "B", fitness: 5},
			&MockIndividual{ID: "C", fitness: 10},
		}

		fittest := fittestIndividual(individuals)
		assert.Equal(t, "C", fittest.ID)
	})

	t.Run("fittest individual is somewhere else in the slice", func(t *testing.T) {
		individuals := []Individual{
			&MockIndividual{ID: "A", fitness: 1},
			&MockIndividual{ID: "B", fitness: 10},
			&MockIndividual{ID: "C", fitness: 5},
		}

		fittest := fittestIndividual(individuals)
		assert.Equal(t, "B", fittest.ID)
	})

	t.Run("there are two fittest individuals and the first is chosen", func(t *testing.T) {
		individuals := []Individual{
			&MockIndividual{ID: "A", fitness: 10},
			&MockIndividual{ID: "B", fitness: 10},
			&MockIndividual{ID: "C", fitness: 5},
		}

		fittest := fittestIndividual(individuals)
		assert.Equal(t, "A", fittest.ID)
	})

	t.Run("zero slice returns nil pointer", func(t *testing.T) {
		var individuals []Individual

		fittest := fittestIndividual(individuals)
		assert.Nil(t, fittest)
	})

	t.Run("singleton slice returns the expected individual", func(t *testing.T) {
		individuals := []Individual{
			&MockIndividual{ID: "A", fitness: 10},
		}

		fittest := fittestIndividual(individuals)
		assert.Equal(t, "A", fittest.ID)
	})
}

func TestSelectForCrossover(t *testing.T) {
	// TODO: test
}