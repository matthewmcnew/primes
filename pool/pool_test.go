package pool_test

import (
	"github.com/matthewmcnew/primes/models"
	"github.com/matthewmcnew/primes/pool"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pool", func() {
	Describe("MaxValue", func() {
		FIt("should find the most common prime divisor", func() {
			pool := pool.NewPool(1)

			eventChan := pool.EventChannel()

			pool.Run(200)

			event1 := <-eventChan
			Expect(event1).To(Equal(&models.ChangeEvent{Prime: 2, Job: 2}))

			event2 := <-eventChan
			Expect(event2).To(Equal(&models.ChangeEvent{Prime: 3, Job: 12}))

			event3 := <-eventChan
			Expect(event3).To(Equal(&models.ChangeEvent{Prime: 5, Job: 80}))

			event4 := <-eventChan
			Expect(event4).To(Equal(&models.ChangeEvent{Prime: 7, Job: 196}))

			Eventually(func() chan *models.ChangeEvent { return eventChan }).Should(BeClosed())
		})
	})
})
