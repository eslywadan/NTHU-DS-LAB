package otelkit

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOtelKit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Otel Kit")
}
