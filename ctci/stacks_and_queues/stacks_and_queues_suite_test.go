package stacks_and_queues_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestStacksAndQueues(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "StacksAndQueues Suite")
}
