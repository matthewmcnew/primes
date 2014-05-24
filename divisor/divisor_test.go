package divisor_test

import (
	"github.com/matthewmcnew/primes/divisor"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Divisor", func() {

	Describe("HigestPrimeDivisor", func() {
		It("should find the higest prime divisor", func() {
			Expect(divisor.HigestPrime(12)).To(Equal(3))
			Expect(divisor.HigestPrime(11)).To(Equal(11))
			Expect(divisor.HigestPrime(30)).To(Equal(5))
			Expect(divisor.HigestPrime(32)).To(Equal(2))
		})
	})
})
