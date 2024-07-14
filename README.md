# jubilant-octo-palm-tree
Generic implementation of a genetic algorithm, in Go.

## Repository structure
This repository contains a package for solving a constraint satisfaction problem. Users of this package should be expected to provide a concrete implementation of `Config`, and can then pass this into `RunGeneticAlgorithm` which abstracts the rest of the implementation.

Some examples of how this can be used may be provided in an `examples/` toplevel directory.

TODO: repository structure in ASCII form

## TODO: how to run

## Why does this exist?
It seemed like a fun thing to implement, and it could be useful for some day-to-day tasks.

## CI/CD pipeline
The repository contains a test suite. The package is unit tested. The examples may also be tested for their validity (TODO).

## When will the binary terminate?
The user is expected to provide an implementation of `ContinuingCondition`, which will halt the program and return the fittest individual when that evaluates to false.

## Input Configuration

The user can supply the following configuration:

* Population size
* Chance of mutation
TODO flesh out

## Todos

* Collect together Todos here
* Add linter config
