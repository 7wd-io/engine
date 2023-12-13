package engine

const (
	Agriculture TokenId = iota + 1
	Architecture
	Economy
	Law
	Masonry
	Mathematics
	Philosophy
	Strategy
	Theology
	Urbanism
)

type TokenId int
type TokenList []TokenId
type TokenMap map[TokenId]Token

type Token struct {
	unit
}
