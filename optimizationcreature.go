package genetic

import (
	"math"
	"math/rand"
)

type OptimizationCreature struct {
	value float32
}

func NewCreature() Creature {
	return &OptimizationCreature{value: rand.Float32()*(2*6.3) - 6.3}
	// return &OptimizationCreature{value: rand.Float32()}
}

func (c *OptimizationCreature) Fitness() float32 {
	x := float64(c.value)
	return float32(math.Sin(x) * x * x)
}

// func (c *OptimizationCreature) Fitness() float32 {
// 	return 0.25 - (c.value-0.5)*(c.value-0.5)
// }

// func (c *OptimizationCreature) Fitness() float32 {
// 	return c.value
// }

func (c *OptimizationCreature) CrossOver(b Creature) (Creature, Creature) {
	aValue, bValue := c.value, b.(*OptimizationCreature).value
	return &OptimizationCreature{value: (aValue*2 + bValue) / 3}, &OptimizationCreature{value: (aValue + bValue*2) / 3}
}

func (c *OptimizationCreature) Mutate() {
	v := c.value
	v += (rand.Float32() - 0.5) / 10
	if v < 0 {
		v = 0
	}
	if v > 1 {
		v = 1
	}
	c.value = v
}
