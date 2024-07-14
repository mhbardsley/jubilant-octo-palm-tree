# jubilant-octo-palm-tree
Generic implementation of a genetic algorithm, in Go.

## Repository structure
This repository contains a package for solving a constraint satisfaction problem. Users of this package should be expected to provide a concrete implementation of `Config`, and can then pass this into `RunGeneticAlgorithm` which abstracts the rest of the implementation.

Some examples of how this can be used may be provided in an `examples/` toplevel directory.

## Why does this exist?
It seemed like a fun thing to implement, and it could be useful for some day-to-day tasks.

## CI/CD pipeline
The repository contains a test suite. The package is unit tested. The examples may also be tested for their validity.

## When will the binary terminate?
The user is expected to provide an implementation of `ContinuingCondition`, which will halt the program and return the fittest individual when that evaluates to false.

## Todos

* Test it with a toy example
* Repo structure in ASCII form
* Concurrency when working with the population
* Some way of asserting that selectForCrossover is given a sorted list
* Some way of asserting the property that the cumulative total will be n(n+1)/2
* selectForCrossover to use binary search mechanism
* See if there is a better way to test the RNG
* Test the examples
* Collect together Todos here
* User guide
* Godoc
