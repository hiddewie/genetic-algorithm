package genetic

import "math/rand"

type Dice struct {
	sides  int
	n      int
	values []int
}

func NewDicePool(size int, selectProportion float32, iterations, n, sides int) *BasicPool {
	p := BasicPool{size: size, selectProportion: selectProportion, iterations: iterations}
	p.pool = make([]Creature, 0, size)
	for i := 0; i < size; i++ {
		p.pool = append(p.pool, NewDice(n, sides))
	}
	return &p
}

func NewDice(n int, sides int) Creature {
	dice := Dice{n: n, sides: sides, values: make([]int, n)}
	for i := 0; i < n; i++ {
		dice.values[i] = rand.Intn(sides)
	}

	return &dice
}

func (c *Dice) Fitness() float32 {
	m := make(map[int]int)
	for i := 0; i < c.n; i++ {
		m[c.values[i]]++
	}

	max := 0
	bestI := 0
	for i := range m {
		if m[i] > max || (m[i] == max && i > bestI) {
			max = m[i]
			bestI = i
		}
	}
	// fmt.Println(fmt.Sprintf("Values %d have fitness %f", c.values, float32(max)+float32(bestI)/float32(c.sides+1)))
	return float32(max) + float32(bestI)/float32(c.sides+1)
}

func (c *Dice) CrossOver(b Creature) (Creature, Creature) {
	d := b.(*Dice)

	index := rand.Intn(c.n)
	values1 := make([]int, c.n)
	copy(values1[index:], c.values[index:])
	copy(values1[:index], d.values[:index])

	values2 := make([]int, c.n)
	copy(values2[index:], d.values[index:])
	copy(values2[:index], c.values[:index])

	// fmt.Println(fmt.Sprintf("Crossing %d and %d to get %d and %d at <%d>", c.values, d.values, values1, values2, index))

	return &Dice{n: c.n, sides: c.sides, values: values1}, &Dice{n: c.n, sides: c.sides, values: values2}
}

func (c *Dice) Mutate() {
	index := rand.Intn(c.n)
	// fmt.Println(fmt.Sprintf("Mutating %d at <%d>", c.values, index))
	c.values[index] = rand.Intn(c.sides)
	// fmt.Println(fmt.Sprintf("Got %d", c.values))
}
