package machinery_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/williammartin/machinery"
)

var _ = Describe("Blueprint", func() {

	It("maintains information about connections", func() {
		start := NewState("start")
		end := NewState("end")
		blueprint := NewBlueprint()
		blueprint.Connect(start, end, "edge")
		node, ok := blueprint.Connection(start, "edge")
		Expect(ok).To(BeTrue())
		Expect(node).To(Equal(end))
	})

	It("returns not ok for non-existent connections", func() {
		blueprint := NewBlueprint()
		_, ok := blueprint.Connection(NewState("start"), "edge")
		Expect(ok).To(BeFalse())
	})

})
