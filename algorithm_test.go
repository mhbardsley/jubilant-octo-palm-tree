package algorithm

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockIndividual struct {
	ID      string
	fitness float64
}

func (m *mockIndividual) Fitness() float64 {
	return m.fitness
}

func (m *mockIndividual) Mutate() {
	// Mock implementation
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

func TestInitializePopulation(t *testing.T) {
	t.Run("Generated population has correct size", func(t *testing.T) {
		populationSize := 10
		generateIndividual := func() *mockIndividual {
			return &mockIndividual{ID: "mock", fitness: 5.0}
		}

		population, _ := initializePopulation(populationSize, generateIndividual)

		assert.Equal(t, populationSize, len(population))
	})

	t.Run("Correctly identifies the best individual", func(t *testing.T) {
		populationSize := 5
		staticID := 1

		generateIndividual := func() *mockIndividual {
			id := staticID
			staticID++
			return &mockIndividual{ID: strconv.Itoa(id), fitness: float64(id)}
		}

		_, bestIndividual := initializePopulation(populationSize, generateIndividual)

		assert.Equal(t, "5", bestIndividual.ID)
	})
}

func TestUpdateBestIndividual(t *testing.T) {
	t.Run("Updates best individual when a better individual is found", func(t *testing.T) {
		// Initial best individual with lower fitness
		bestIndividual := &mockIndividual{ID: "1", fitness: 5.0}

		// Population with an individual that has better fitness
		population := []*mockIndividual{
			{ID: "2", fitness: 7.0},
			{ID: "3", fitness: 6.0},
		}

		updatedBestIndividual := updateBestIndividual(population, bestIndividual)

		// Assert that the best individual is updated to the one with the highest fitness
		assert.Equal(t, "2", updatedBestIndividual.ID)
	})

	t.Run("Keeps current best individual when no better individual is found", func(t *testing.T) {
		// Initial best individual with higher fitness
		bestIndividual := &mockIndividual{ID: "1", fitness: 8.0}

		// Population with individuals that have lower fitness
		population := []*mockIndividual{
			{ID: "2", fitness: 7.0},
			{ID: "3", fitness: 6.0},
		}

		updatedBestIndividual := updateBestIndividual(population, bestIndividual)

		// Assert that the best individual remains the same
		assert.Equal(t, "1", updatedBestIndividual.ID)
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
