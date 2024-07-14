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
	fixedValue int
}

func (m *mockRng) Intn(_ int) int {
	return m.fixedValue
}

func TestGeneratePopulation(_ *testing.T) {
	// TODO: implement
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
	t.Run("first item is the one to be selected", func(t *testing.T) {
		individuals := []*mockIndividual{
			&mockIndividual{ID: "A"},
			&mockIndividual{ID: "B"},
			&mockIndividual{ID: "C"},
		}
		rng := &mockRng{0}

		selected := selectForCrossoverRng(individuals, rng)
		assert.Equal(t, "A", selected.ID)
	})
	t.Run("last item is the one to be selected with the lower part of the range", func(t *testing.T) {
		individuals := []*mockIndividual{
			&mockIndividual{ID: "A"},
			&mockIndividual{ID: "B"},
			&mockIndividual{ID: "C"},
		}
		rng := &mockRng{3}

		selected := selectForCrossoverRng(individuals, rng)
		assert.Equal(t, "C", selected.ID)
	})
	t.Run("last item is the one to be selected with the upper part of the range", func(t *testing.T) {
		individuals := []*mockIndividual{
			&mockIndividual{ID: "A"},
			&mockIndividual{ID: "B"},
			&mockIndividual{ID: "C"},
		}
		rng := &mockRng{5}

		selected := selectForCrossoverRng(individuals, rng)
		assert.Equal(t, "C", selected.ID)
	})
	t.Run("middle item is the one to be selected with the lower part of the range", func(t *testing.T) {
		individuals := []*mockIndividual{
			&mockIndividual{ID: "A"},
			&mockIndividual{ID: "B"},
			&mockIndividual{ID: "C"},
		}
		rng := &mockRng{1}

		selected := selectForCrossoverRng(individuals, rng)
		assert.Equal(t, "B", selected.ID)
	})
	t.Run("middle item is the one to be selected with the upper part of the range", func(t *testing.T) {
		individuals := []*mockIndividual{
			&mockIndividual{ID: "A"},
			&mockIndividual{ID: "B"},
			&mockIndividual{ID: "C"},
		}
		rng := &mockRng{2}

		selected := selectForCrossoverRng(individuals, rng)
		assert.Equal(t, "B", selected.ID)
	})
	t.Run("singleton list", func(t *testing.T) {
		individuals := []*mockIndividual{
			&mockIndividual{ID: "A"},
		}
		rng := &mockRng{0}

		selected := selectForCrossoverRng(individuals, rng)
		assert.Equal(t, "A", selected.ID)
	})
}
