package worker

import (
	"github.com/matthewmcnew/primes/divisor"
	"github.com/matthewmcnew/primes/models"
)

type tallyManagerInterface interface {
	NewValue(value *models.CalculatedResult)
}

func Work(inputChan chan int, tallyManager tallyManagerInterface) {
	var highestPrime int
	for {
		job, ok := <-inputChan
		if !ok {
			return
		}

		highestPrime = divisor.HigestPrime(job)
		tallyManager.NewValue(&models.CalculatedResult{Prime: highestPrime, Job: job})
	}
}
