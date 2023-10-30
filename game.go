package main

type Mutator interface {
	mutate(s *State) error
}

type mutator func(s *State)

type burner interface {
	burn(s *State)
}

type scorable interface {
	getPoints(s *State) int
}
