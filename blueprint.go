package machinery

type connection map[string]string

type Blueprint struct {
	connections map[string]connection
	states      map[string]*State
}

func NewBlueprint() *Blueprint {
	return &Blueprint{
		connections: make(map[string]connection),
		states:      make(map[string]*State),
	}
}

func (b *Blueprint) Connect(startState, endState *State, trigger string) {
	connections, ok := b.connections[startState.Label]
	if !ok {
		b.connections[startState.Label] = make(map[string]string)
		b.Connect(startState, endState, trigger)
		return
	}

	connections[trigger] = endState.Label
	b.states[startState.Label] = startState
	b.states[endState.Label] = endState
}

func (b *Blueprint) Connection(state *State, trigger string) (*State, bool) {
	connections, ok := b.connections[state.Label]
	if !ok {
		return nil, false
	}

	node, ok := connections[trigger]
	if !ok {
		return nil, false
	}

	return b.states[node], true
}
