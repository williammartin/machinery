package machinery_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/williammartin/machinery"
)

var _ = Describe("A Phone Call", func() {

	It("represents a phone call", func() {
		callCount := 0
		callInProgress := false

		offHook := machinery.NewState("OffHook")
		dialing := machinery.NewState("Dialing")
		connected := machinery.NewState("Connected")
		connected.OnEntry(func() error {
			callCount++
			callInProgress = true
			return nil
		})
		connected.OnExit(func() error {
			callInProgress = false
			return nil
		})

		blueprint := machinery.NewBlueprint()
		blueprint.Connect(offHook, dialing, "Dial")
		blueprint.Connect(dialing, connected, "Connect")
		blueprint.Connect(connected, offHook, "Hangup")

		machine := machinery.NewMachine(blueprint, offHook)
		Expect(machine.State()).To(Equal(offHook))

		machine.Fire("Dial")
		Expect(machine.State()).To(Equal(dialing))

		machine.Fire("Connect")
		Expect(machine.State()).To(Equal(connected))
		Expect(callCount).To(Equal(1))
		Expect(callInProgress).To(BeTrue())

		machine.Fire("Hangup")
		Expect(machine.State()).To(Equal(offHook))
		Expect(callInProgress).To(BeFalse())
	})
})
