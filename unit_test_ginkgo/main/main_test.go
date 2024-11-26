package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"unit_test_ginkgo/main"
)

var _ = Describe("Main", func() {
	Describe("LuasBelahKetupat", func() {
		It("should calculate the area correctly for positive diagonal values", func() {
			diagonal1 := 10.0
			diagonal2 := 8.0
			expected := 40.0

			result := main.LuasBelahKetupat(diagonal1, diagonal2)

			Expect(result).To(Equal(expected))
		})

		It("should return 0 if one of the diagonals is 0", func() {
			diagonal1 := 0.0
			diagonal2 := 8.0
			expected := 0.0

			result := main.LuasBelahKetupat(diagonal1, diagonal2)

			Expect(result).To(Equal(expected))
		})

		It("should return 0 if both diagonals are 0", func() {
			diagonal1 := 0.0
			diagonal2 := 0.0
			expected := 0.0

			result := main.LuasBelahKetupat(diagonal1, diagonal2)

			Expect(result).To(Equal(expected))
		})
	})
})
