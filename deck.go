package engine

import (
	"strings"
)

func newDeck(cards CardList, layout string) *deck {
	d := &deck{
		graph:    make(graph),
		cards:    cards,
		faceDown: cardSet{},
		layout:   layout,
	}

	layout = strings.TrimLeft(layout, "\n")

	prev := make(line)
	curr := make(line)

	inRowPos := 0
	rowN := 1
	var cardPos int

	for _, char := range layout {
		switch char {
		case '\n':
			for pos, cid := range curr {
				d.graph[cid] = make([]CardId, 2)

				if prev[pos+1] > 0 {
					d.graph[prev[pos+1]][0] = cid
				}

				if prev[pos-1] > 0 {
					d.graph[prev[pos-1]][1] = cid
				}
			}
			inRowPos = 0
			rowN++
			prev = curr
			curr = make(line)
		case '[':
			if rowN%2 == 0 {
				d.faceDown[d.cards[cardPos]] = struct{}{}
			}

			curr[inRowPos] = d.cards[cardPos]
			cardPos++
			fallthrough
		default:
			inRowPos++
		}
	}

	return d
}

type deck struct {
	graph    graph
	cards    CardList
	faceDown cardSet
	layout   string
}

// getLayout format: [...value,]
// value == -2: guild slot
// value == -1: empty slot
// value == 0 face down card
// value > 0 face up card id
func (dst *deck) getLayout() CardList {
	cards := make(CardList, len(dst.cards))
	copy(cards, dst.cards)

	for i, cid := range dst.cards {
		if _, ok := dst.graph[cid]; !ok {
			cards[i] = -1
			continue
		}

		if _, ok := dst.faceDown[cid]; ok {
			if R.Cards[cid].Group != groupGuild {
				cards[i] = 0
			} else {
				cards[i] = -2
			}

			continue
		}
	}

	return cards
}

func (dst *deck) playableCards() cardSet {
	cards := make(cardSet)

	for parent, children := range dst.graph {
		if children[0] == 0 && children[1] == 0 {
			cards[parent] = struct{}{}
		}
	}

	return cards
}

func (dst *deck) topLineCards() CardList {
	var cards CardList
	count := strings.Count(strings.Split(strings.Trim(dst.layout, "\n"), "\n")[0], "[")

	for _, cid := range dst.getLayout()[:count] {
		if cid > 0 {
			cards = append(cards, cid)
		}
	}

	return cards
}

func (dst *deck) returnedCards() CardList {
	search := make(cardSet, len(dst.cards))

	for _, cid := range dst.cards {
		search[cid] = struct{}{}
	}

	var cards CardList
	age := R.Cards[dst.cards[0]].Age

	for _, cid := range R.ageCards[age] {
		if _, ok := search[cid]; !ok {
			cards = append(cards, cid)
		}
	}

	return cards
}

func (dst *deck) removeCard(cid CardId) {
	delete(dst.graph, cid)

	for parent, children := range dst.graph {
		if children[0] == cid {
			children[0] = 0
		}

		if children[1] == cid {
			children[1] = 0
		}

		if children[0] == 0 && children[1] == 0 {
			delete(dst.faceDown, parent)
		}
	}
}

func (dst *deck) isEmpty() bool {
	return len(dst.graph) == 0
}

type graph map[CardId]CardList

type line map[int]CardId
