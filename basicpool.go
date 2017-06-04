package genetic

import "sort"

type BasicPool struct {
	size             int
	pool             []Creature
	selectProportion float32
	iterations       int
}

func NewPool(size int, selectProportion float32, iterations int) *BasicPool {
	p := BasicPool{size: size, selectProportion: selectProportion, iterations: iterations}
	p.pool = make([]Creature, 0, size)
	for i := 0; i < size; i++ {
		p.pool = append(p.pool, NewCreature())
	}
	return &p
}

func (p *BasicPool) Size() int {
	return len(p.pool)
}

func (p *BasicPool) PoolSize() int {
	return p.size
}

func (p *BasicPool) Get(index int) Creature {
	return p.pool[index]
}

func (p *BasicPool) Finished(iteration int, maxFitness, lastMaxFitness, averageFitness, lastAverageFitness float32, time int64) bool {
	return iteration >= p.iterations
}

func (p *BasicPool) Select() []Creature {
	sort.Slice(p.pool, func(i, j int) bool { return p.pool[i].Fitness() < p.pool[j].Fitness() })
	startIndex := int((1.0 - p.selectProportion) * float32(len(p.pool)))
	return p.pool[startIndex:]
}

func (p *BasicPool) MutationProbability() float32 {
	return 0.1
}

func (p *BasicPool) SetCreatures(creatures []Creature) {
	p.pool = creatures
}
