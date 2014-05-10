package goraph_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoraph(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goraph Suite")
}
