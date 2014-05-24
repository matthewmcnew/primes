package pool

import (
	"github.com/matthewmcnew/primes/tally"
	"github.com/matthewmcnew/primes/worker"

	"sync"
)

type Pool struct {
	cpus int
}

func NewPool(numCPUs int) *Pool {
	return &Pool{cpus: numCPUs}
}

func (p *Pool) Run(x int) int {
	inputChan := make(chan int, p.cpus)
	tallyManager := tally.NewTallyManager()

	job := &job{pool: p, inputChan: inputChan, tallyManager: tallyManager}
	job.Add(p.cpus)

	job.RunJob(x)

	job.Wait()

	mostCommon, _ := job.tallyManager.MostCommon()
	return mostCommon
}

type job struct {
	pool         *Pool
	inputChan    chan int
	tallyManager *tally.TallyManager
	sync.WaitGroup
}

func (j *job) RunJob(x int) {
	go func() {
		for i := 2; i <= x; i++ {
			j.inputChan <- i
		}
		close(j.inputChan)
	}()

	for i := 1; i <= j.pool.cpus; i++ {
		go func() {
			worker.Work(j.inputChan, j.tallyManager)
			j.Done()
		}()
	}
}
