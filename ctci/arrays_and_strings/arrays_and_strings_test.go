package arrays_and_strings_test

import (
	. "github.com/floreks/training/ctci/arrays_and_strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ArraysAndStrings", func() {
	It("Test IsUnique", func() {
		Expect(IsUnique("asd")).To(Equal(true))
		Expect(IsUnique("assd")).To(Equal(false))
	})

	It("Test IsUniqueRaw", func() {
		Expect(IsUniqueRaw("asd")).To(Equal(true))
		Expect(IsUniqueRaw("sads")).To(Equal(false))
	})

	It("Test IsPermutation", func() {
		Expect(IsPermutation("asd", "das")).To(Equal(true))
		Expect(IsPermutation("asd", "dasf")).To(Equal(false))
		Expect(IsPermutation("ass", "das")).To(Equal(false))
	})

	It("Test URLify", func() {
		Expect(URLify("Mr John Smith    ", 13)).To(Equal("Mr%20John%20Smith"))
	})

	It("Test IsPalindromePermutation", func() {
		Expect(IsPalindromePermutation("Tact Coa")).To(Equal(true))
		Expect(IsPalindromePermutation("test")).To(Equal(false))
		Expect(IsPalindromePermutation("a")).To(Equal(true))
		Expect(IsPalindromePermutation("aaaaaaa")).To(Equal(true))
		Expect(IsPalindromePermutation("bbabb")).To(Equal(true))
		Expect(IsPalindromePermutation("aaabbaaa")).To(Equal(true))
		Expect(IsPalindromePermutation("aaabbbaaa")).To(Equal(true))
		Expect(IsPalindromePermutation("")).To(Equal(true))
	})

	It("Test IsPalindromePermutationBit", func() {
		Expect(IsPalindromePermutationBit("Tact Coa")).To(Equal(true))
		Expect(IsPalindromePermutationBit("test")).To(Equal(false))
		Expect(IsPalindromePermutationBit("a")).To(Equal(true))
		Expect(IsPalindromePermutationBit("aaaaaaa")).To(Equal(true))
		Expect(IsPalindromePermutationBit("bbabb")).To(Equal(true))
		Expect(IsPalindromePermutationBit("aaabbaaa")).To(Equal(true))
		Expect(IsPalindromePermutationBit("aaabbbaaa")).To(Equal(true))
		Expect(IsPalindromePermutationBit("")).To(Equal(true))
	})

	It("Test IsOneAway", func() {
		Expect(IsOneAway("aaaa", "aa")).To(Equal(false))
		Expect(IsOneAway("aaaa", "aaa")).To(Equal(true))
		Expect(IsOneAway("aaaaa", "aab")).To(Equal(false))
		Expect(IsOneAway("aaaa", "aab")).To(Equal(false))
		Expect(IsOneAway("aaaa", "aaab")).To(Equal(true))
		Expect(IsOneAway("pale", "ple")).To(Equal(true))
		Expect(IsOneAway("pales", "pale")).To(Equal(true))
		Expect(IsOneAway("pale", "bale")).To(Equal(true))
		Expect(IsOneAway("pale", "bake")).To(Equal(false))
		Expect(IsOneAway("pale", "elapa")).To(Equal(true))
	})

	It("Test IsOneAwaySimple", func() {
		Expect(IsOneAwaySimple("aaaa", "aa")).To(Equal(false))
		Expect(IsOneAwaySimple("aaa", "aaba")).To(Equal(true))
		Expect(IsOneAwaySimple("aaaa", "aaa")).To(Equal(true))
		Expect(IsOneAwaySimple("aaaa", "aaab")).To(Equal(true))
		Expect(IsOneAwaySimple("pale", "ple")).To(Equal(true))
		Expect(IsOneAwaySimple("pafe", "ple")).To(Equal(false))
		Expect(IsOneAwaySimple("pales", "pale")).To(Equal(true))
		Expect(IsOneAwaySimple("pale", "bale")).To(Equal(true))
		Expect(IsOneAwaySimple("pale", "paff")).To(Equal(false))
		Expect(IsOneAwaySimple("pale", "bake")).To(Equal(false))
	})

	It("Test Compress", func() {
		Expect(Compress("aabcccccaaa")).To(Equal("a2b1c5a3"))
		Expect(Compress("aabcccccaaab")).To(Equal("a2b1c5a3b1"))
		Expect(Compress("abcd")).To(Equal("abcd"))
		Expect(Compress("aabbcd")).To(Equal("aabbcd"))
		Expect(Compress("aabbccddd")).To(Equal("a2b2c2d3"))
	})

	It("Test Rotate", func() {
		Expect(Rotate([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})).To(Equal([][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}))
		Expect(Rotate([][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}})).To(Equal([][]int{{13, 9, 5, 1}, {14, 10, 6, 2}, {15, 11, 7, 3}, {16, 12, 8, 4}}))
	})

	It("Test ZeroMatrix", func() {
		Expect(ZeroMatrix([][]int{{1, 0, 3}, {4, 5, 6}, {7, 8, 9}})).To(Equal([][]int{{0, 0, 0}, {4, 0, 6}, {7, 0, 9}}))
	})

	It("Test IsRotation", func() {
		Expect(IsRotation("waterbottle", "erbottlewat")).To(Equal(true))
		Expect(IsRotation("waterbottle", "erbottlewar")).To(Equal(false))
		Expect(IsRotation("asd", "aaaaa")).To(Equal(false))
		Expect(IsRotation("asd", "dsa")).To(Equal(false))
	})
})
