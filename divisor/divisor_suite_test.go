package divisor_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDivisor(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Divisor Suite")
}
