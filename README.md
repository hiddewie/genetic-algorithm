This repository is a small library written in Go, suitable for the basic running
of a genetic algorithm (see `genetic.go`).

The data consists of a pool of creatures (`Pool`). The pool controls the managing of the creatures (`Creature`).
The creatures control the behavior: fitness, mutation and crossover.

For the `Pool` a basic implementation `BasicPool` exists (see `basicpool.go`). This generates a pool and manages the creatures.

For the `Creature`s two basic implementations exist. The first implementation optimizes a function and is called `OptimizationCreature` (see `optimizationcreature.go`). The fitness function is the function value to be maximized. The mutation is a small increase or decrease of the value. The crossover is the changing of values of two creatures to a value closer to their average.

The second `Creature` is a `Dice` (see `dice.go`). The creature holds `n` dice with values in the range `[0, sides)`. The fitness function is the number of equal dice, where a mutation randomly assigns one die. The crossover function interchanges a random part of the dice with another creature.

The tests (`*_test.go`) runs an algorithm for both optimization examples.
