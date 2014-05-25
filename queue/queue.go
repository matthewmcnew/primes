package queue

import (
	"github.com/matthewmcnew/primes/models"
)

// A Queue implements heap.Interface and holds Items.
type Queue []*models.CalculatedResult

func (pq Queue) Len() int { return len(pq) }

func (pq Queue) Less(i, j int) bool {
    return pq[i].Job < pq[j].Job
}

func (pq Queue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
}

func (pq *Queue) Push(x interface{}) {
    item := x.(*models.CalculatedResult)
    *pq = append(*pq, item)
}

func (pq *Queue) Pop() interface{} {
     old := *pq
     n := len(old)
     item := old[n-1]
     *pq = old[0 : n-1]
     return item
}

func (pq Queue) LowestJob() interface{} {
    return pq[0].Job
}