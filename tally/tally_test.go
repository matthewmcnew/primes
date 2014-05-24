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
						tallyManager.NewValue(2)
						tallyManager.NewValue(5)
						tallyManager.NewValue(2)
						tallyManager.NewValue(2)

						highest,value := tallyManager.MostCommon()
						Expect(highest).To(Equal(2))
						Expect(value).To(Equal(3))
				})
			})
})
