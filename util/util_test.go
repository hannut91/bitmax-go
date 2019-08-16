package util

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Util", func() {
	Describe("UUID", func() {
		It("returns uuid", func() {
			uuid := UUID()

			Expect(uuid).To(HaveLen(32))
		})
	})

	Describe("ComputeHmac256", func() {
		encoded := "i19IcCmVwVmMVz2x4hhmqbgl1KeU0WnXBgoDYFeWNgs="

		It("returns hmac256", func() {
			hmac256 := ComputeHmac256("message", "secret")

			Expect(hmac256).To(Equal(encoded))
		})
	})
})
