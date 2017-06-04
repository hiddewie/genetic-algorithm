package genetic

import (
	"math/rand"
	"testing"
	"time"
)

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

var _ Creature = (*Dice)(nil)

func TestRunDicePool(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	pool := NewDicePool(100, 0.4, 5, 6, 10)
	Run(pool)
}
