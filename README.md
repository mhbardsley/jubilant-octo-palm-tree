# jubilant-octo-palm-tree
Generic implementation of a genetic algorithm, in Go.

## Repository structure
This repository contains a package for solving a constraint satisfaction problem. Users of this package provide an `Individual` implementation along with `GenerateIndividual`, `Crossover`, and `ContinuingCondition` callbacks via `Config`, then pass it to `RunGeneticAlgorithm`.

See `examples/` for end-to-end usage.

## Why does this exist?
It seemed like a fun thing to implement, and it could be useful for some day-to-day tasks.

## CI/CD pipeline
The repository contains a test suite. The package is unit tested. The examples may also be tested for their validity.

## When will the binary terminate?
The user is expected to provide an implementation of `ContinuingCondition`, which will halt the program and return the fittest individual when that evaluates to false.

## Todos

* Repo structure in ASCII form
* Test the examples
* User guide
* Godoc
