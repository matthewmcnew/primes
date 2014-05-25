package tally

import (
	"container/heap"
	"github.com/matthewmcnew/primes/models"
	"github.com/matthewmcnew/primes/queue"
)

type TallyManager struct {
	Counts          map[int]int
	Events          chan *models.ChangeEvent
	inputChan       chan *models.CalculatedResult
	mostCommonValue int
	queue           *queue.Queue
	next            int
}

func NewTallyManager(staringPoint int) *TallyManager {
	counts := make(map[int]int)
	events := make(chan *models.ChangeEvent, 100)
	inputChan := make(chan *models.CalculatedResult, 100)

	queue := &queue.Queue{}
	heap.Init(queue)

	return &TallyManager{Counts: counts, Events: events, inputChan: inputChan, next: staringPoint, queue: queue}
}

func (t *TallyManager) NewValue(calculatedResult *models.CalculatedResult) {
	t.inputChan <- calculatedResult
}

func (t *TallyManager) Run() {
	for calculatedResult := range t.inputChan {
		t.ensureOrder(calculatedResult)
	}
	close(t.Events)
}

func (t *TallyManager) Close() {
	close(t.inputChan)
}

func (t *TallyManager) ensureOrder(calculatedResult *models.CalculatedResult) {
	if calculatedResult.Job != t.next {
		heap.Push(t.queue, calculatedResult)
	} else {
		t.next += 1
		t.tallyResult(calculatedResult)
		if t.queue.Len() > 0 && t.queue.LowestJob() == t.next {
			nextCalculatedResult := heap.Pop(t.queue).(*models.CalculatedResult)
			t.ensureOrder(nextCalculatedResult)
		}
	}
}

func (t *TallyManager) tallyResult(calculatedResult *models.CalculatedResult) {
	t.Counts[calculatedResult.Prime] += 1

	if t.Counts[calculatedResult.Prime] > t.Counts[t.mostCommonValue] {
		t.mostCommonValue = calculatedResult.Prime
		t.Events <- &models.ChangeEvent{Prime: calculatedResult.Prime, Job: calculatedResult.Job}
	}
}
