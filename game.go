package engine

type Mutator interface {
	Mutate(s *State) error
}

type mutator func(s *State)

type Burner interface {
	Burn(s *State)
}

type Scorable interface {
	GetPoints(s *State) int
}
