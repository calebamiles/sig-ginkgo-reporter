package reporter_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestSigGinkgoReporter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SigGinkgoReporter Suite")
}
