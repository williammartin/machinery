package machinery_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/williammartin/machinery"
)

var _ = Describe("Machine", func() {

	It("returns the starting state after creation", func() {
		machine := NewMachine(nil, NewState("start"))
		Expect(machine.State()).To(Equal(NewState("start")))
	})

	Context("when provided connections", func() {
		It("can change to one of those states", func() {
			blueprint := NewBlueprint()
			start := NewState("start")
			end := NewState("end")
			blueprint.Connect(start, end, "trigger")

			machine := NewMachine(blueprint, start)
			Expect(machine.Fire("trigger")).To(Succeed())
			Expect(machine.State()).To(Equal(end))
		})
	})

	Context("when a trigger isn't defined", func() {
		It("errors and doesn't change state", func() {
			start := NewState("start")
			machine := NewMachine(NewBlueprint(), start)
			Expect(machine.Fire("trigger")).To(MatchError("Trigger 'trigger' is not defined"))
			Expect(machine.State()).To(Equal(start))
		})
	})

	Context("when a state has a registered onEntry function", func() {
		It("is called", func() {
			blueprint := NewBlueprint()
			start := NewState("start")
			end := NewState("end")

			ended := false
			end.OnEntry(func() error {
				ended = true
				return nil
			})
			blueprint.Connect(start, end, "trigger")

			machine := NewMachine(blueprint, start)
			Expect(machine.Fire("trigger")).To(Succeed())
			Expect(ended).To(BeTrue())
		})

		Context("but the function errors", func() {
			It("propagates that error but the transition still occurs", func() {
				blueprint := NewBlueprint()
				start := NewState("start")
				end := NewState("end")

				end.OnEntry(func() error {
					return errors.New("EXPLODED")
				})
				blueprint.Connect(start, end, "trigger")

				machine := NewMachine(blueprint, start)
				Expect(machine.Fire("trigger")).To(MatchError("EXPLODED"))
				Expect(machine.State()).To(Equal(end))
			})
		})
	})

	Context("when a state has a registered onExit function", func() {
		It("is called", func() {
			blueprint := NewBlueprint()
			start := NewState("start")
			end := NewState("end")

			transitioned := false
			start.OnExit(func() error {
				transitioned = true
				return nil
			})
			blueprint.Connect(start, end, "trigger")

			machine := NewMachine(blueprint, start)
			Expect(machine.Fire("trigger")).To(Succeed())
			Expect(transitioned).To(BeTrue())
		})

		Context("but the function errors", func() {
			It("propagates that error and the transition doesn't occur", func() {
				blueprint := NewBlueprint()
				start := NewState("start")
				end := NewState("end")

				start.OnExit(func() error {
					return errors.New("EXPLODED")
				})
				blueprint.Connect(start, end, "trigger")

				machine := NewMachine(blueprint, start)
				Expect(machine.Fire("trigger")).To(MatchError("EXPLODED"))
				Expect(machine.State()).To(Equal(start))
			})
		})
	})
})
