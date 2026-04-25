package algorithm

import (
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
		selected := tournamentSelectRng(population, &mockRng{results: []int{1, 3}})

		assert.Equal(t, "D", selected.ID)
	})

	t.Run("returns the same individual when both picks collide", func(t *testing.T) {
		selected := tournamentSelectRng(population, &mockRng{results: []int{2, 2}})

		assert.Equal(t, "C", selected.ID)
	})

	t.Run("returns zero value for an empty population", func(t *testing.T) {
		var empty []*mockIndividual

		assert.Nil(t, tournamentSelectRng(empty, &mockRng{}))
	})
}
