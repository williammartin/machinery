package machinery

type State struct {
	Label       string
	onEntryFunc func() error
	onExitFunc  func() error
}

func NewState(label string) *State {
	return &State{
		Label: label,
	}
}

func (s *State) OnEntry(f func() error) {
	s.onEntryFunc = f
}

func (s *State) Entered() error {
	if s.onEntryFunc == nil {
		return nil
	}
	return s.onEntryFunc()
}

func (s *State) OnExit(f func() error) {
	s.onExitFunc = f
}

func (s *State) Exited() error {
	if s.onExitFunc == nil {
		return nil
	}
	return s.onExitFunc()
}
