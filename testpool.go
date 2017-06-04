package genetic

import (
	"math"
	"math/rand"
	"sort"
)

type TestPool struct {
	size             int
	pool             []Creature
	selectProportion float32
	iterations       int
}
type TestCreature struct {
	value float32
}

func NewPool(size int, selectProportion float32, iterations int) *TestPool {
	p := TestPool{size: size, selectProportion: selectProportion, iterations: iterations}
	p.pool = make([]Creature, 0, size)
	for i := 0; i < size; i++ {
		p.pool = append(p.pool, NewCreature())
	}
	return &p
}

func (p *TestPool) Size() int {
	return len(p.pool)
}

func (p *TestPool) PoolSize() int {
	return p.size
}

func (p *TestPool) Get(index int) Creature {
	return p.pool[index]
}

func NewCreature() Creature {
	return &TestCreature{value: rand.Float32()*(2*6.3) - 6.3}
	// return &TestCreature{value: rand.Float32()}
}

func (p *TestPool) Finished(iteration int, maxFitness, lastMaxFitness, averageFitness, lastAverageFitness float32, time int64) bool {
	return iteration >= p.iterations
}

func (p *TestPool) Select() []Creature {
	sort.Slice(p.pool, func(i, j int) bool { return p.pool[i].Fitness() < p.pool[j].Fitness() })
	startIndex := int((1.0 - p.selectProportion) * float32(len(p.pool)))
	return p.pool[startIndex:]
}

func (c *TestCreature) Fitness() float32 {
	x := float64(c.value)
	return float32(math.Sin(x) * x * x)
}

// func (c *TestCreature) Fitness() float32 {
// 	return 0.25 - (c.value-0.5)*(c.value-0.5)
// }

// func (c *TestCreature) Fitness() float32 {
// 	return c.value
// }

func (c *TestCreature) CrossOver(b Creature) (Creature, Creature) {
	aValue, bValue := c.value, b.(*TestCreature).value
	return &TestCreature{value: (aValue*2 + bValue) / 3}, &TestCreature{value: (aValue + bValue*2) / 3}
}

func (c *TestCreature) Mutate() {
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

func (p *TestPool) MutationProbability() float32 {
	return 0.1
}

func (p *TestPool) SetCreatures(creatures []Creature) {
	p.pool = creatures
}
