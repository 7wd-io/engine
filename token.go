package main

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
type tokenList []TokenId
type tokenMap map[TokenId]*token

type token struct {
	unit
}
