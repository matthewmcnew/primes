package pool

import (
	"github.com/matthewmcnew/primes/tally"
	"github.com/matthewmcnew/primes/worker"
	"github.com/matthewmcnew/primes/models"

	"sync"
)

type Pool struct {
	cpus int
	tallyManager *tally.TallyManager
}

func NewPool(numCPUs int) *Pool {
	tallyManager := tally.NewTallyManager(2)
	return &Pool{cpus: numCPUs, tallyManager: tallyManager}
}

func (p *Pool) EventChannel() chan *models.ChangeEvent {
	return p.tallyManager.Events
}

func (p *Pool) Run(x int) {
	go p.tallyManager.Run()

	inputChan := make(chan int, p.cpus)

	job := &job{pool: p, inputChan: inputChan}
	job.Add(p.cpus)

	job.RunJob(x)

	job.Wait()
	p.tallyManager.Close()
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
			worker.Work(j.inputChan, j.pool.tallyManager)
			j.Done()
		}()
	}
}
