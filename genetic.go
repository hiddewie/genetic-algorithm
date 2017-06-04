package genetic

import (
	"fmt"
	"math/rand"
	"time"
)

type Creature interface {
	Fitness() float32
}

type GeneticPool interface {
	NewCreature() Creature
	Get(index int) Creature
	Size() int
	PoolSize() int
	// Add(...Creature)
	SetCreatures([]Creature)
	Finished(iteration int, maxFitness, lastMaxFitness, averageFitness, lastAverageFitness float32, time int64) bool
	Select() []Creature
	CrossOver(Creature, Creature) (Creature, Creature)
	Mutate(index int)
	MutationProbability() float32
}

func MaxFitness(p GeneticPool) (Creature, float32) {
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

func AverageFitness(p GeneticPool) float32 {
	var sum float32
	for i := 0; i < p.Size(); i++ {
		sum += p.Get(i).Fitness()
	}
	return sum / float32(p.Size())
}

func WeightedSelect(breeding []Creature, maxFitness float32) Creature {
	for {
		// Only select from the breeding stock
		index := rand.Intn(len(breeding))
		if rand.Float32() < breeding[index].Fitness()/maxFitness {
			return breeding[index]
		}
	}
}

func Run(p GeneticPool) Creature {
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

		// Fill pool: Crossover
		bestCreature, maxFitness = MaxFitness(p)
		averageFitness = AverageFitness(p)
		for len(breeding) < p.Size() {
			first := WeightedSelect(breeding, maxFitness)
			other := WeightedSelect(breeding, maxFitness)
			for {
				if first != other {
					break
				}
				other = WeightedSelect(breeding, maxFitness)
			}
			a, b := p.CrossOver(first, other)
			breeding = append(breeding, a, b)
		}
		p.SetCreatures(breeding)

		// Mutate
		for i := 0; i < p.Size(); i++ {
			if rand.Float32() < p.MutationProbability() {
				p.Mutate(i)
			}
		}

		t = (time.Now().UnixNano() - startTime) / (int64(time.Millisecond) / int64(time.Nanosecond))
		fmt.Println(fmt.Sprintf("Iteration %d [%d ms], max fitness %f, average fitness %f.", iteration, t, maxFitness, averageFitness))
		iteration++
	}

	fmt.Println(fmt.Sprintf("Finished after %d ms. Best fitness is %f.", t, maxFitness))

	return bestCreature
}
