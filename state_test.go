package machinery_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/williammartin/machinery"
)

var _ = Describe("State", func() {

	It("calls the registered onEntry function when entered", func() {
		state := NewState("state")
		callCount := 0
		state.OnEntry(func() error {
			callCount++
			return nil
		})

		state.Entered()
		Expect(callCount).To(Equal(1))
	})

	It("returns errors from registered onEntry functions", func() {
		state := NewState("state")
		state.OnEntry(func() error {
			return errors.New("EXPLODED")
		})

		Expect(state.Entered()).To(MatchError("EXPLODED"))
	})

	It("calls the registered onExit function when exited", func() {
		state := NewState("state")
		callCount := 0
		state.OnExit(func() error {
			callCount++
			return nil
		})

		state.Exited()
		Expect(callCount).To(Equal(1))
	})

	It("returns errors from registered onExit functions", func() {
		state := NewState("state")
		state.OnExit(func() error {
			return errors.New("EXPLODED")
		})

		Expect(state.Exited()).To(MatchError("EXPLODED"))
	})
})
