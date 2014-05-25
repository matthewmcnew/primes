package worker_test

import (
	"github.com/matthewmcnew/primes/worker"
	"github.com/matthewmcnew/primes/models"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type fakeTallyManager struct {
	values []*models.CalculatedResult
}

func (t *fakeTallyManager) NewValue(value *models.CalculatedResult) {
	t.values = append(t.values, value)
}

var _ = Describe("Worker", func() {
	Describe("Work", func() {
		It("should record the most common HigestPrime divisors", func() {
			inputChan := make(chan int, 2)
			tallyManager := &fakeTallyManager{values: make([]*models.CalculatedResult, 0)}

			inputChan <- 12
			inputChan <- 30
			close(inputChan)

			worker.Work(inputChan, tallyManager)

			Expect(tallyManager.values[0]).To(Equal(&models.CalculatedResult{Prime: 3, Job: 12}))
			Expect(tallyManager.values[1]).To(Equal(&models.CalculatedResult{Prime: 5, Job: 30}))
		})
	})
})
