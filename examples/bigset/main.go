// Run a genetic algorithm to gather a fixed-size set with the biggest total value
package main

import (
	"fmt"
	algo "github.com/mhbardsley/jubilant-octo-palm-tree"
	"math/rand"
	"time"
)

func main() {
	config := algo.Config[*individual]{PopulationSize: 1000, AlgorithmConfig: &algorithm{time.Now()}}
	fmt.Println(algo.RunGeneticAlgorithm(config))
}

type individual struct {
	set     []int
	fitness *float64
}

var numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

func (i *individual) Fitness() float64 {
	if i.fitness != nil {
		return *(i.fitness)
	}
	sum := 0.0
	for _, val := range i.set {
		sum += float64(val)
	}
	i.fitness = &sum
	return sum
}

func (i *individual) Mutate() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	duplicate := make([]int, len(i.set))
	copy(duplicate, i.set)
	for x := range duplicate {
		if rng.Float64() < 0.1 {
			duplicate[x] = numbers[rand.Intn(len(numbers))]
		}
	}
	newInd := &individual{set: duplicate}
	if newInd.Fitness() > i.Fitness() {
		*i = *newInd
	}
}

type algorithm struct {
	startTime time.Time
}

func (a *algorithm) ContinuingCondition() bool {
	return time.Now().Sub(a.startTime) <= 1*time.Second
}

func (a *algorithm) GenerateIndividual() *individual {
	set := make([]int, 10)
	sum := 0.0
	for x := range 10 {
		set[x] = numbers[rand.Intn(len(numbers))]
		sum += float64(set[x])
	}
	return &individual{set, &sum}
}

func (a *algorithm) GenerateCrossover(ind1, ind2 *individual) *individual {
	return &individual{append(ind1.set[:5], ind2.set[5:]...), nil}
}
