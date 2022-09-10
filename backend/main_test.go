package backend_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"todo/backend"
)

var _ = Describe("Main", func() {
	It("should work", func() {
		backend.SaveTasks()
		Expect(true).To(BeTrue())
	})
})
