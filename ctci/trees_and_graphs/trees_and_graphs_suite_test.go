package trees_and_graphs_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTreesAndGraphs(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TreesAndGraphs Suite")
}
