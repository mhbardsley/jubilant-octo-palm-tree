package main

import (
	"fmt"

	"github.com/mhbardsley/jubilant-octo-palm-tree/pkg"
)

func main() {
	config := pkg.Config{}
	_ = pkg.RunGeneticAlgorithm(config)
	// load in the configuration
	fmt.Println("TODO: processing")
	// run algorithm
	fmt.Println("TODO: handled output")
	// TODO: output individual
}