package main

import (
	"fmt"
)

func main() {
	// load in the configuration
	config := pkg.LoadConfiguration()
	fmt.Println("TODO: load in data")
	data := pkg.LoadData()
	fmt.Println("TODO: processing")
	selected := pkg.RunGeneticAlgorithm(config, data)  // TODO: is parameter passing the best way to do this?
	fmt.Println("TODO: handled output")
	fmt.Println("The selected individual is %s", string(selected)) // TODO: some sort of stringify function
}