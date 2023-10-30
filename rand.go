package main

import (
	"math/rand"
	"time"
)

func randBool() bool {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(2) != 0
}

func shuffleTokens(in tokenList) tokenList {
	out := make(tokenList, len(in))
	copy(out, in)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(out), func(i, j int) {
		out[i], out[j] = out[j], out[i]
	})

	return out
}

func shuffleWonders(in wonderList) wonderList {
	out := make(wonderList, len(in))
	copy(out, in)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(out), func(i, j int) {
		out[i], out[j] = out[j], out[i]
	})

	return out
}

func randCards(a age) cardList {
	var out cardList
	cards := shuffleCards(R.ageCards[a])

	switch a {
	case ageIII:
		out = cards[:deckLimit-guildsLimit]
		out = append(out, shuffleCards(R.guilds)[:guildsLimit]...)
		out = shuffleCards(out)
	default:
		out = cards[:deckLimit]
	}

	return out
}

func shuffleCards(in cardList) cardList {
	out := make(cardList, len(in))
	copy(out, in)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(out), func(i, j int) {
		out[i], out[j] = out[j], out[i]
	})

	return out
}
