package in_memory_platform_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestInMemoryPlatform(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "InMemoryPlatform Suite")
}
