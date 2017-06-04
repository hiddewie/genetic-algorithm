package genetic

import (
	"fmt"
	"math/rand"
	"time"
)

// Creature defines an element of a pool for the genetic algorithm
type Creature interface {
	Fitness() float32
}

// Pool defines a pool containing creatures and for running a genetic algorithm
type Pool interface {
	NewCreature() Creature
	Get(index int) Creature
	Size() int
	PoolSize() int
	SetCreatures([]Creature)
	Finished(iteration int, maxFitness, lastMaxFitness, averageFitness, lastAverageFitness float32, time int64) bool
	Select() []Creature
	CrossOver(Creature, Creature) (Creature, Creature)
	Mutate(index int)
	MutationProbability() float32
}

// MaxFitness detemines the maximum fitness and the creature in the pool
func MaxFitness(p Pool) (Creature, float32) {
	var max float32
	var best Creature
	for i := 0; i < p.Size(); i++ {
		cur := p.Get(i).Fitness()
		if cur > max {
			max = cur
			best = p.Get(i)
		}
	}
	return best, max
}

// AverageFitness detemines the average fitness of the pool
func AverageFitness(p Pool) float32 {
	var sum float32
	for i := 0; i < p.Size(); i++ {
		sum += p.Get(i).Fitness()
	}
	return sum / float32(p.Size())
}

// WeightedSelect selects creatures from the breeding stock weighted by fitness
func WeightedSelect(breeding []Creature, maxFitness float32) Creature {
	for {
		// Only select from the breeding stock
		index := rand.Intn(len(breeding))
		if rand.Float32() < breeding[index].Fitness()/maxFitness {
			return breeding[index]
		}
	}
}

// Run runs the genetic pool and finds the best creature
func Run(p Pool) Creature {
	startTime := time.Now().UnixNano()
	var maxFitness, lastMaxFitness, averageFitness, lastAverageFitness float32
	var t int64
	iteration := 0
	var bestCreature Creature

	fmt.Println(fmt.Sprintf("Running pool of size %d.", p.PoolSize()))

	for !(p.Finished(iteration, maxFitness, lastMaxFitness, averageFitness, lastAverageFitness, t)) {
		lastMaxFitness = maxFitness
		lastAverageFitness = averageFitness

		// Select breeding
		breeding := p.Select()

		// New pool
		newPool := make([]Creature, 0, p.PoolSize())
		newPool = append(newPool, breeding...)

		// Fill pool: Crossover
		bestCreature, maxFitness = MaxFitness(p)
		averageFitness = AverageFitness(p)
		for len(newPool) < p.Size() {
			// Select creatures for breeding
			first := WeightedSelect(breeding, maxFitness)
			other := WeightedSelect(breeding, maxFitness)
			for {
				if first != other {
					break
				}
				other = WeightedSelect(breeding, maxFitness)
			}
			// Crossover creatures
			a, b := p.CrossOver(first, other)
			newPool = append(newPool, a, b)
		}
		// Update the creatures
		p.SetCreatures(newPool)

		// Mutate
		for i := 0; i < p.Size(); i++ {
			if rand.Float32() < p.MutationProbability() {
				p.Mutate(i)
			}
		}

		t = (time.Now().UnixNano() - startTime) / (int64(time.Millisecond) / int64(time.Nanosecond))
		iteration++
		fmt.Println(fmt.Sprintf("Iteration %d [%d ms, (av. %d ms)], max fitness %f, average fitness %f.", iteration, t, t/int64(iteration), maxFitness, averageFitness))
	}

	fmt.Println(fmt.Sprintf("Finished after %d ms. Best fitness is %f. Best creature is %s", t, maxFitness, bestCreature))

	return bestCreature
}
