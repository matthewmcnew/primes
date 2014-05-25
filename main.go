package main

import (
	"github.com/matthewmcnew/primes/pool"

	"flag"
	"fmt"
	"runtime"
	"strconv"
)

func main() {
	flag.Parse()

	numCPUS, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Println("Please Provide Valid # of Num Cpus")
		return
	}

	maximumNumberToComputeTo, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		fmt.Println("Please Provide Valid maximum number to convert to")
		return
	}

	runtime.GOMAXPROCS(numCPUS)

	pool := pool.NewPool(numCPUS)
	pool.Run(maximumNumberToComputeTo)

	for event := range pool.EventChannel() {
		fmt.Printf("Now %d is the most Common Starting at %d.\n", event.Prime, event.Job)
	}
}
