package pool_test

import (
	"github.com/matthewmcnew/primes/pool"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Pool", func() {
	Describe("MaxValue", func() {
		It("should find the most common prime divisor", func() {
			pool := pool.NewPool(1)

			Expect(pool.Run(2)).To(Equal(2))
			Expect(pool.Run(12)).To(Equal(3))
			Expect(pool.Run(80)).To(Equal(5))
			Expect(pool.Run(196)).To(Equal(7))
		})
	})
})
