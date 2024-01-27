package engine

import (
	"math/rand"
	"time"
)

func randBool() bool {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(2) != 0
}

func shuffleTokens(in TokenList) TokenList {
	out := make(TokenList, len(in))
	copy(out, in)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(out), func(i, j int) {
		out[i], out[j] = out[j], out[i]
	})

	return out
}

func shuffleWonders(in WonderList) WonderList {
	out := make(WonderList, len(in))
	copy(out, in)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(out), func(i, j int) {
		out[i], out[j] = out[j], out[i]
	})

	return out
}

func randCards(a Age) CardList {
	var out CardList
	cards := shuffleCards(R.ageCards[a])

	switch a {
	case AgeIII:
		out = cards[:deckLimit-guildsLimit]
		out = append(out, shuffleCards(R.guilds)[:guildsLimit]...)
		out = shuffleCards(out)
	default:
		out = cards[:deckLimit]
	}

	return out
}

func shuffleCards(in CardList) CardList {
	out := make(CardList, len(in))
	copy(out, in)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(out), func(i, j int) {
		out[i], out[j] = out[j], out[i]
	})

	return out
}
