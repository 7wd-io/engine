package engine

const (
	phaseNil phase = iota
	phasePrepare
	PhaseTurn
	phaseSelectWhoBeginsTheNextAge
	phasePickBoardToken
	phasePickRandomToken
	phaseBurnCard
	phasePickDiscardedCard
	phasePickTopLineCard
	phasePickReturnedCards
	phaseOver
)

type phase int
