package engine

type unit struct {
	Effects []interface{} `json:"effects"`
}

func (dst unit) Mutate(s *State) error {
	for _, v := range dst.Effects {
		eff, ok := v.(Mutator)

		if !ok {
			continue
		}

		if err := eff.Mutate(s); err != nil {
			return err
		}
	}

	return nil
}

func (dst unit) Burn(s *State) {
	for _, v := range dst.Effects {
		eff, ok := v.(Burner)

		if !ok {
			continue
		}

		eff.Burn(s)
	}
}

func (dst unit) GetPoints(s *State) int {
	var points int

	for _, v := range dst.Effects {
		eff, ok := v.(Scorable)

		if !ok {
			continue
		}

		points += eff.GetPoints(s)
	}

	return points
}
