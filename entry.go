package engine

func StateFrom(m ...Mutator) (*State, error) {
	s := new(State)

	for _, v := range m {
		if err := v.Mutate(s); err != nil {
			return s, err
		}
	}

	return s, nil
}
