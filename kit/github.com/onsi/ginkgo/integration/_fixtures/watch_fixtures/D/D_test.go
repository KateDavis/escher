package D_test

import (
	. "github.com/gocircuit/escher/kit/github.com/onsi/ginkgo/integration/_fixtures/watch_fixtures/C"

	. "github.com/gocircuit/escher/kit/github.com/onsi/ginkgo"
	. "github.com/gocircuit/escher/kit/github.com/onsi/gomega"
)

var _ = Describe("D", func() {
	It("should do it", func() {
		Ω(DoIt()).Should(Equal("done!"))
	})
})
