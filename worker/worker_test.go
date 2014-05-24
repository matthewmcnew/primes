package worker_test

import (
	"github.com/matthewmcnew/primes/worker"
	//	"github.com/matthewmcnew/primes/tally"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type fakeTallyManager struct {
	values []int
}

func (t *fakeTallyManager) NewValue(value int) {
	t.values = append(t.values, value)
}

var _ = Describe("Worker", func() {
	Describe("Work", func() {
		It("should record the most common HigestPrime divisors", func() {
			inputChan := make(chan int, 2)
			tallyManager := &fakeTallyManager{values: make([]int, 0)}

			inputChan <- 12
			inputChan <- 30
			close(inputChan)

			worker.Work(inputChan, tallyManager)

			Expect(tallyManager.values[0]).To(Equal(3))
			Expect(tallyManager.values[1]).To(Equal(5))
		})
	})
})
