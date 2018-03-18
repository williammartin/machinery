package machinery

import (
	"fmt"
)

type Machine struct {
	blueprint *Blueprint
	state     *State
}

func NewMachine(blueprint *Blueprint, startState *State) *Machine {
	return &Machine{
		blueprint: blueprint,
		state:     startState,
	}
}

func (m *Machine) State() *State {
	return m.state
}

func (m *Machine) Fire(trigger string) error {
	state, ok := m.blueprint.Connection(m.state, trigger)
	if !ok {
		return fmt.Errorf("Trigger '%s' is not defined", trigger)
	}

	if err := m.state.Exited(); err != nil {
		return err
	}

	m.state = state
	return m.state.Entered()
}
