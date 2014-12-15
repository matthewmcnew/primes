## Introduction

## Setup

Make sure you have go installed and your $GOPATH is properly set. Then run:

 go get -u github.com/matthewmcnew/primes

## Run

go install github.com/matthewmcnew/primes

This will install an executable named primes to $GOPATH/bin

## Run the tests

 We used Ginkgo for our tests.

 go get github.com/onsi/ginkgo/ginkgo

 go get github.com/onsi/gomega

 #add $GOPATH/bin to your path if necessary

 ginkgo -r src/github.com/matthewmcnew/primes
