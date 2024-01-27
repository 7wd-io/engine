package engine

import (
	"errors"
	"fmt"
)

var (
	ErrActionNotAllowed = errors.New("action not allowed")
	ErrStateFrom        = func(index int, m Mutator, wrap error) error {
		return fmt.Errorf("state from failed by move.index=%d (%#v) (%w)", index, m, wrap)
	}
)
