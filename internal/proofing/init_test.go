package proofing_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestProofing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "internal/proofing")
}
