package engine

const (
	Civilian Victory = iota + 1
	MilitarySupremacy
	ScienceSupremacy
	Resign
	Timeout
)

type Victory int
