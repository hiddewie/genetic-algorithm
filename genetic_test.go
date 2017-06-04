package genetic

import (
	"math/rand"
	"sort"
	"testing"
	"time"
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
		p.pool = append(p.pool, p.NewCreature())
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

func (p *TestPool) NewCreature() Creature {
	return &TestCreature{value: rand.Float32()}
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
	return 0.25 - (c.value-0.5)*(c.value-0.5)
}

// func (c *TestCreature) Fitness() float32 {
// 	return c.value
// }

func (p *TestPool) CrossOver(a, b Creature) (Creature, Creature) {
	aValue, bValue := a.(*TestCreature).value, b.(*TestCreature).value
	return &TestCreature{value: (aValue*2 + bValue) / 3}, &TestCreature{value: (aValue + bValue*2) / 3}
}

func (p *TestPool) Mutate(index int) {
	v := p.pool[index].Fitness()
	v += (rand.Float32() - 0.5) / 10
	if v < 0 {
		v = 0
	}
	if v > 1 {
		v = 1
	}
	p.pool[index].(*TestCreature).value = v
}

func (p *TestPool) MutationProbability() float32 {
	return 0.1
}

func (p *TestPool) SetCreatures(creatures []Creature) {
	p.pool = creatures
}

var _ Pool = (*TestPool)(nil)
var _ Creature = (*TestCreature)(nil)

func TestTestPool(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	pool := NewPool(100, 0.4, 3)
	if pool.Size() != 100 {
		t.Fail()
	}
}

func TestRun(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	pool := NewPool(100, 0.4, 5)
	Run(pool)
}
