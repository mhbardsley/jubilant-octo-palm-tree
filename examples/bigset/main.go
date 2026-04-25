// Run a genetic algorithm to gather a fixed-size set with the biggest total value.
package main

import (
	"fmt"
	"math/rand"
	"time"

	algo "github.com/mhbardsley/jubilant-octo-palm-tree"
)

const (
	setSize        = 10
	populationSize = 1000
	mutationRate   = 0.1
	runDuration    = time.Second
)

var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

type individual struct {
	set []int
}

func (i *individual) Fitness() float64 {
	sum := 0.0
	for _, v := range i.set {
		sum += float64(v)
	}
	return sum
}

func (i *individual) Mutate() {
	for x := range i.set {
		if rand.Float64() < mutationRate {
			i.set[x] = numbers[rand.Intn(len(numbers))]
		}
	}
}

func generate() *individual {
	set := make([]int, setSize)
	for x := range set {
		set[x] = numbers[rand.Intn(len(numbers))]
	}
	return &individual{set: set}
}

func crossover(a, b *individual) *individual {
	split := setSize / 2
	child := make([]int, setSize)
	copy(child[:split], a.set[:split])
	copy(child[split:], b.set[split:])
	return &individual{set: child}
}

func main() {
	deadline := time.Now().Add(runDuration)
	best := algo.RunGeneticAlgorithm(algo.Config[*individual]{
		PopulationSize:      populationSize,
		GenerateIndividual:  generate,
		Crossover:           crossover,
		ContinuingCondition: func() bool { return time.Now().Before(deadline) },
	})
	fmt.Println(best.set, best.Fitness())
}
