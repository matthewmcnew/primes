package tally_test

import (
	"github.com/matthewmcnew/primes/models"
	"github.com/matthewmcnew/primes/tally"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tally", func() {

	Describe("MostCommon", func() {
		Context("In Order", func() {
			It("Fires events when a new Most Common item is reached", func() {
				tallyManager := tally.NewTallyManager(1)

				go tallyManager.Run()

				first := &models.CalculatedResult{Prime: 1, Job: 1}
				tallyManager.NewValue(first)

				event := <-tallyManager.Events
				Expect(event).To(Equal(&models.ChangeEvent{Prime: 1, Job: 1}))

				tallyManager.NewValue(&models.CalculatedResult{Prime: 1, Job: 2})
				tallyManager.NewValue(&models.CalculatedResult{Prime: 5, Job: 3})
				tallyManager.NewValue(&models.CalculatedResult{Prime: 5, Job: 4})

				leadTaker := &models.CalculatedResult{Prime: 5, Job: 5}
				tallyManager.NewValue(leadTaker)

				event = <-tallyManager.Events
				Expect(event).To(Equal(&models.ChangeEvent{Prime: 5, Job: 5}))
			})
		})

		Context("Out of Order", func() {
			It("Fires events when a new Most Common item is reached", func() {
				tallyManager := tally.NewTallyManager(100)

				go tallyManager.Run()

				first := &models.CalculatedResult{Prime: 1, Job: 100}
				tallyManager.NewValue(first)

				event := <-tallyManager.Events
				Expect(event).To(Equal(&models.ChangeEvent{Prime: 1, Job: 100}))

				tallyManager.NewValue(&models.CalculatedResult{Prime: 5, Job: 124})

				leadTaker := &models.CalculatedResult{Prime: 2, Job: 102}
				tallyManager.NewValue(leadTaker)

				tallyManager.NewValue(&models.CalculatedResult{Prime: 2, Job: 101})

				event = <-tallyManager.Events
				Expect(event).To(Equal(&models.ChangeEvent{Prime: 2, Job: 102}))
			})
		})

	})

	Describe("Close", func() {
		It("Closes the Input Chan", func() {
			tallyManager := tally.NewTallyManager(1)

			go tallyManager.Run()

			tallyManager.Close()

			Eventually(func() chan *models.ChangeEvent { return tallyManager.Events }).Should(BeClosed())
			Expect(func() { tallyManager.NewValue(&models.CalculatedResult{Prime: 5, Job: 124}) }).To(Panic())
		})
	})
})
