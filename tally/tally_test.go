package tally_test

import (
	"github.com/matthewmcnew/primes/tally"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tally", func() {

	Describe("MostCommon", func() {
		It("should find the higest prime divisor", func() {
			tallyManager := tally.NewTallyManager()
			for i := 1; i < 10; i++ {
				tallyManager.NewValue(i)
			}
			tallyManager.NewValue(2)
			tallyManager.NewValue(2)
			tallyManager.NewValue(2)

			highest, value := tallyManager.MostCommon()
			Expect(highest).To(Equal(2))
			Expect(value).To(Equal(4))
		})
	})
})
