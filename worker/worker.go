package worker

import (
	"github.com/matthewmcnew/primes/divisor"
)

type tallyManagerInterface interface {
	NewValue(value int)
}

func Work(inputChan chan int, tallyManager tallyManagerInterface) {
	for {
		job, ok := <-inputChan
		if !ok {
			return
		}

		tallyManager.NewValue(divisor.HigestPrime(job))
	}
}
