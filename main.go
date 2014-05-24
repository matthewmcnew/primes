package main

import (
	"github.com/matthewmcnew/primes/pool"

	"flag"
	"strconv"
	"fmt"
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

	pool := pool.NewPool(numCPUS)

	result := pool.Run(maximumNumberToComputeTo)

	fmt.Println("Most Common:", result)
}
