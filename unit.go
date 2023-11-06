package engine

type unit struct {
	Effects []interface{} `json:"effects"`
}

func (dst unit) mutate(s *State) error {
	for _, v := range dst.Effects {
		eff, ok := v.(Mutator)

		if !ok {
			continue
		}

		if err := eff.mutate(s); err != nil {
			return err
		}
	}

	return nil
}

func (dst unit) burn(s *State) {
	for _, v := range dst.Effects {
		eff, ok := v.(burner)

		if !ok {
			continue
		}

		eff.burn(s)
	}
}

func (dst unit) getPoints(s *State) int {
	var points int

	for _, v := range dst.Effects {
		eff, ok := v.(scorable)

		if !ok {
			continue
		}

		points += eff.getPoints(s)
	}

	return points
}
