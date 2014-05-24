package tally

import (
	"sync"
)

type TallyManager struct {
	Counts map[int]int
	sync.Mutex
}

func NewTallyManager() *TallyManager {
	counts := make(map[int]int)
	return &TallyManager{Counts: counts}
}

func (t *TallyManager) NewValue(value int) {
	t.Lock()
	defer t.Unlock()
	t.Counts[value] = t.Counts[value] + 1
}

func (t *TallyManager) MostCommon() (mostCommon, maximum int) {
	t.Lock()
	defer t.Unlock()

	mostCommon = 0
	maximum = 0
	for key, value := range t.Counts {
		if value > maximum {
			maximum = value
			mostCommon = key
		}
	}

	return
}
