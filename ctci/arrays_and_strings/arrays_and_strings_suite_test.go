package arrays_and_strings_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestArraysAndStrings(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ArraysAndStrings Suite")
}
