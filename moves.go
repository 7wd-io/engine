package main

import (
	"encoding/json"
	"errors"
	"log"
)

const (
	mPrepare moveId = iota + 1
	mPickWonder
	mPickBoardToken
	mConstructCard
	mConstructWonder
	mDiscardCard
	mSelectWhoBeginsTheNextAge
	mBurnCard
	mPickRandomToken
	mPickTopLineCard
	mPickDiscardedCard
	mPickReturnedCards
	mOver
)

type moveId int

type move struct {
	Id moveId `json:"id"`
}

func NewMovePrepare(p1, p2 Nickname) Mutator {
	if randBool() {
		p2, p1 = p1, p2
	}

	t := shuffleTokens(R.tokenIds)

	w := shuffleWonders(R.wonderIds)

	c := map[age]cardList{}

	for _, a := range []age{ageI, ageII, ageIII} {
		c[a] = randCards(a)
	}

	return prepareMove{
		move:         move{mPrepare},
		P1:           p1,
		P2:           p2,
		Wonders:      w[:wonderSelectionPoolSize*2],
		Tokens:       t[:startingTokensCount],
		RandomTokens: t[startingTokensCount : startingTokensCount+randomTokensCount],
		Cards:        c,
	}
}

type prepareMove struct {
	move
	P1           Nickname         `json:"p1"`
	P2           Nickname         `json:"p2"`
	Wonders      wonderList       `json:"wonders"`
	Tokens       tokenList        `json:"tokens"`
	RandomTokens tokenList        `json:"randomTokens"`
	Cards        map[age]cardList `json:"cards"`
}

func (dst prepareMove) mutate(s *State) error {
	if s.Phase != phaseNil {
		return ErrActionNotAllowed
	}

	s.Age = ageI
	s.Phase = phasePrepare
	s.FirstTurn = dst.P1

	s.Tokens = make(tokenList, len(dst.Tokens))
	copy(s.Tokens, dst.Tokens)

	s.Me = newCity(dst.P1)
	s.Enemy = newCity(dst.P2)

	s.Dialogs = []mutator{}
	s.RandomItems.Cards = dst.Cards
	s.RandomItems.Tokens = dst.RandomTokens
	s.RandomItems.Wonders = dst.Wonders
	s.DialogItems.Wonders = make(wonderList, wonderSelectionPoolSize)
	copy(s.DialogItems.Wonders, dst.Wonders[:wonderSelectionPoolSize])

	return nil
}

func NewMovePickWonder(w WonderId) Mutator {
	return pickWonderMove{
		move:   move{mPickWonder},
		Wonder: w,
	}
}

type pickWonderMove struct {
	move
	Wonder WonderId `json:"wonder"`
}

func (dst pickWonderMove) mutate(s *State) error {
	if s.Phase != phasePrepare {
		return ErrActionNotAllowed
	}

	foundWonder := false
	// pull used wonder
	for i, w := range s.DialogItems.Wonders {
		if w == dst.Wonder {
			s.DialogItems.Wonders[i] = 0
			foundWonder = true
			break
		}
	}

	if !foundWonder {
		return ErrActionNotAllowed
	}

	s.Me.Wonders.add(dst.Wonder)

	countPicked := s.Me.Wonders.countTotal() + s.Enemy.Wonders.countTotal()

	// pick scheme
	// [N] - player
	// stage 1: [1][2][2][1]
	// stage 2: [2][1][1][2]
	// after first move 1
	switch countPicked {
	case 2, 6:
		//  2 wonders in a row
	default:
		s.nextTurn()
		// normal flow, next player
	}

	switch countPicked {
	case wonderSelectionPoolSize:
		copy(s.DialogItems.Wonders, s.RandomItems.Wonders[wonderSelectionPoolSize:])
	case wonderSelectionPoolSize * 2:
		s.Phase = PhaseTurn
		s.DialogItems.Wonders = nil

		s.Deck = newDeck(s.RandomItems.Cards[ageI], R.layouts[ageI])
		s.refreshDeckCards()
		s.refreshCities()
	}

	return nil
}

func NewMoveSelectWhoBeginsTheNextAge(n Nickname) Mutator {
	return selectWhoBeginsTheNextAgeMove{
		move:   move{mSelectWhoBeginsTheNextAge},
		Player: n,
	}
}

type selectWhoBeginsTheNextAgeMove struct {
	move
	Player Nickname `json:"player"`
}

func (dst selectWhoBeginsTheNextAgeMove) mutate(s *State) error {
	if s.Phase != phaseSelectWhoBeginsTheNextAge {
		return ErrActionNotAllowed
	}

	s.setTurn(dst.Player)

	s.Phase = PhaseTurn

	return nil
}

func NewMovePickBoardToken(t TokenId) Mutator {
	return pickBoardTokenMove{
		move:  move{mPickBoardToken},
		Token: t,
	}
}

type pickBoardTokenMove struct {
	move
	Token TokenId `json:"token"`
}

func (dst pickBoardTokenMove) mutate(s *State) error {
	if s.Phase != phasePickBoardToken {
		return ErrActionNotAllowed
	}

	var found bool

	for _, tid := range s.DialogItems.Tokens {
		if dst.Token == tid {
			found = true
			break
		}
	}

	if !found {
		return ErrActionNotAllowed
	}

	s.Me.Tokens.add(dst.Token)

	if err := R.tokens[dst.Token].mutate(s); err != nil {
		return err
	}

	var origInd int

	for i, item := range s.Tokens {
		if item == dst.Token {
			origInd = i
			break
		}
	}

	// just set 0 to save order tokens
	// on client zero handled as no token
	s.Tokens[origInd] = 0

	s.after()

	return nil
}

func NewMovePickRandomToken(t TokenId) Mutator {
	return pickRandomTokenMove{
		move:  move{mPickRandomToken},
		Token: t,
	}
}

type pickRandomTokenMove struct {
	move
	Token TokenId `json:"token"`
}

func (dst pickRandomTokenMove) mutate(s *State) error {
	if s.Phase != phasePickRandomToken {
		return ErrActionNotAllowed
	}

	var found bool

	for _, tid := range s.DialogItems.Tokens {
		if dst.Token == tid {
			found = true
			break
		}
	}

	if !found {
		return ErrActionNotAllowed
	}

	s.Me.Tokens.add(dst.Token)

	if err := R.tokens[dst.Token].mutate(s); err != nil {
		return err
	}

	s.after()

	return nil
}

func NewMoveConstructCard(c CardId) Mutator {
	return constructCardMove{
		move: move{mConstructCard},
		Card: c,
	}
}

type constructCardMove struct {
	move
	Card CardId `json:"card"`
}

func (dst constructCardMove) mutate(s *State) error {
	if s.Phase != PhaseTurn {
		return ErrActionNotAllowed
	}

	if !s.CardItems.isPlayable(dst.Card) {
		return ErrActionNotAllowed
	}

	if s.Me.Chains.has(dst.Card) {
		if s.Me.Tokens.has(Urbanism) {
			s.Me.Treasure.add(4)
		}
	} else {
		if err := s.pay(getDiscountContextByCard(dst.Card), R.Cards[dst.Card].Cost); err != nil {
			return err
		}
	}

	s.Me.Cards.add(dst.Card)

	s.Deck.removeCard(dst.Card)

	if err := R.Cards[dst.Card].mutate(s); err != nil {
		return err
	}

	s.after()

	return nil
}

func NewMoveConstructWonder(w WonderId, c CardId) Mutator {
	return constructWonderMove{
		move:   move{mConstructWonder},
		Wonder: w,
		Card:   c,
	}
}

type constructWonderMove struct {
	move
	Wonder WonderId `json:"wonder"`
	Card   CardId   `json:"card"`
}

func (dst constructWonderMove) mutate(s *State) error {
	if s.Phase != PhaseTurn {
		return ErrActionNotAllowed
	}

	if !s.CardItems.isPlayable(dst.Card) {
		return ErrActionNotAllowed
	}

	if !s.Me.Wonders.has(dst.Wonder) {
		return ErrActionNotAllowed
	}

	if s.Me.Wonders.isConstructed(dst.Wonder) {
		return ErrActionNotAllowed
	}

	if err := s.pay(dcWonders, R.Wonders[dst.Wonder].Cost); err != nil {
		return err
	}

	s.Deck.removeCard(dst.Card)

	s.Me.Wonders.construct(dst.Wonder, dst.Card)

	totalConstructed := s.Me.Wonders.countConstructed() + s.Enemy.Wonders.countConstructed()

	if totalConstructed == wondersConstructLimit {
		s.Me.Wonders.removeNotConstructed()
		s.Enemy.Wonders.removeNotConstructed()
	}

	if err := R.Wonders[dst.Wonder].mutate(s); err != nil {
		return err
	}

	if s.Me.Tokens.has(Theology) {
		s.PlayAgain = true
	}

	s.after()

	return nil
}

func NewMoveDiscardCard(c CardId) Mutator {
	return discardCardMove{
		move: move{mDiscardCard},
		Card: c,
	}
}

type discardCardMove struct {
	move
	Card CardId `json:"card"`
}

func (dst discardCardMove) mutate(s *State) error {
	if s.Phase != PhaseTurn {
		return ErrActionNotAllowed
	}

	if !s.CardItems.isPlayable(dst.Card) {
		return ErrActionNotAllowed
	}

	s.CardItems.addDiscarded(dst.Card)
	s.Deck.removeCard(dst.Card)
	s.Me.Treasure.add(s.Me.Bank.DiscardReward)

	s.after()

	return nil
}

func NewMoveBurnCard(c CardId) Mutator {
	return burnCardMove{
		move: move{mBurnCard},
		Card: c,
	}
}

type burnCardMove struct {
	move
	Card CardId `json:"card"`
}

func (dst burnCardMove) mutate(s *State) error {
	if s.Phase != phaseBurnCard {
		return ErrActionNotAllowed
	}

	var found bool

	for _, cid := range s.DialogItems.Cards {
		if dst.Card == cid {
			found = true
			break
		}
	}

	if !found {
		return ErrActionNotAllowed
	}

	s.Enemy.Cards.remove(dst.Card)

	R.Cards[dst.Card].burn(s)

	s.after()

	return nil
}

func NewMovePickTopLineCard(c CardId) Mutator {
	return pickTopLineCardMove{
		move: move{mPickTopLineCard},
		Card: c,
	}
}

type pickTopLineCardMove struct {
	move
	Card CardId `json:"card"`
}

func (dst pickTopLineCardMove) mutate(s *State) error {
	if s.Phase != phasePickTopLineCard {
		return ErrActionNotAllowed
	}

	var found bool

	for _, cid := range s.DialogItems.Cards {
		if dst.Card == cid {
			found = true
			break
		}
	}

	if !found {
		return ErrActionNotAllowed
	}

	s.Me.Cards.add(dst.Card)

	s.Deck.removeCard(dst.Card)

	if err := R.Cards[dst.Card].mutate(s); err != nil {
		return err
	}

	s.after()

	return nil
}

func NewMovePickDiscardedCard(c CardId) Mutator {
	return pickDiscardedCardMove{
		move: move{mPickDiscardedCard},
		Card: c,
	}
}

type pickDiscardedCardMove struct {
	move
	Card CardId `json:"card"`
}

func (dst pickDiscardedCardMove) mutate(s *State) error {
	if s.Phase != phasePickDiscardedCard {
		return ErrActionNotAllowed
	}

	var found bool

	for _, cid := range s.DialogItems.Cards {
		if dst.Card == cid {
			found = true
			break
		}
	}

	if !found {
		return ErrActionNotAllowed
	}

	s.Me.Cards.add(dst.Card)
	s.CardItems.removeDiscarded(dst.Card)

	if err := R.Cards[dst.Card].mutate(s); err != nil {
		return err
	}

	s.after()

	return nil
}

func NewMovePickReturnedCards(p CardId, g CardId) Mutator {
	return pickReturnedCardsMove{
		move: move{mPickReturnedCards},
		Pick: p,
		Give: g,
	}
}

type pickReturnedCardsMove struct {
	move
	Pick CardId `json:"pick"`
	Give CardId `json:"give"`
}

func (dst pickReturnedCardsMove) mutate(s *State) error {
	if s.Phase != phasePickReturnedCards {
		return ErrActionNotAllowed
	}

	if !dst.isValidCards(s) {
		return ErrActionNotAllowed
	}

	s.byCity(s.Me.Name, func() {
		s.Me.Cards.add(dst.Pick)
		_ = R.Cards[dst.Pick].mutate(s)
	})

	s.byCity(s.Enemy.Name, func() {
		s.Me.Cards.add(dst.Give)
		_ = R.Cards[dst.Give].mutate(s)
	})

	s.after()

	return nil
}

func (dst pickReturnedCardsMove) isValidCards(s *State) bool {
	var pickOk, giveOk bool

	if dst.Pick == dst.Give {
		return false
	}

	for _, cid := range s.DialogItems.Cards {
		if cid == dst.Pick {
			pickOk = true
		}

		if cid == dst.Give {
			giveOk = true
		}

		if pickOk && giveOk {
			return true
		}
	}

	return false
}

func NewMoveOver(loser Nickname, reason Victory) Mutator {
	return overMove{
		move:   move{mOver},
		Loser:  loser,
		Reason: reason,
	}
}

type overMove struct {
	move
	Loser  Nickname `json:"loser"`
	Reason Victory  `json:"reason"`
}

func (dst overMove) mutate(s *State) error {
	s.over(dst.Reason, s.enemyFor(dst.Loser))

	return nil
}

type Moves []Mutator

func (dst *Moves) UnmarshalJSON(bytes []byte) error {
	var messages []*json.RawMessage

	if err := json.Unmarshal(bytes, &messages); err != nil {
		panic("moves unmarshal fail")
	}

	var m map[string]interface{}
	v := make(Moves, len(messages))

	for index, message := range messages {
		if err := json.Unmarshal(*message, &m); err != nil {
			log.Fatalln(err)
		}

		switch moveId(m["id"].(float64)) {
		case mPrepare:
			var m1 prepareMove

			if err := json.Unmarshal(*message, &m1); err != nil {
				panic("moves unmarshal fail")
			}

			v[index] = m1
		case mPickWonder:
			var m2 pickWonderMove

			if err := json.Unmarshal(*message, &m2); err != nil {
				panic("moves unmarshal fail")
			}

			v[index] = m2
		case mPickBoardToken:
			var m3 pickBoardTokenMove

			if err := json.Unmarshal(*message, &m3); err != nil {
				panic("moves unmarshal fail")
			}

			v[index] = m3
		case mConstructCard:
			var m4 constructCardMove

			if err := json.Unmarshal(*message, &m4); err != nil {
				panic("moves unmarshal fail")
			}

			v[index] = m4
		case mConstructWonder:
			var m5 constructWonderMove

			if err := json.Unmarshal(*message, &m5); err != nil {
				panic("moves unmarshal fail")
			}

			v[index] = m5
		case mDiscardCard:
			var m6 discardCardMove

			if err := json.Unmarshal(*message, &m6); err != nil {
				panic("moves unmarshal fail")
			}

			v[index] = m6
		case mSelectWhoBeginsTheNextAge:
			var m7 selectWhoBeginsTheNextAgeMove

			if err := json.Unmarshal(*message, &m7); err != nil {
				panic("moves unmarshal fail")
			}

			v[index] = m7
		case mBurnCard:
			var m8 burnCardMove

			if err := json.Unmarshal(*message, &m8); err != nil {
				panic("moves unmarshal fail")
			}

			v[index] = m8
		case mPickRandomToken:
			var m9 pickRandomTokenMove

			if err := json.Unmarshal(*message, &m9); err != nil {
				panic("moves unmarshal fail")
			}

			v[index] = m9
		case mPickTopLineCard:
			var m10 pickTopLineCardMove

			if err := json.Unmarshal(*message, &m10); err != nil {
				panic("moves unmarshal fail")
			}

			v[index] = m10
		case mPickDiscardedCard:
			var m11 pickDiscardedCardMove

			if err := json.Unmarshal(*message, &m11); err != nil {
				panic("moves unmarshal fail")
			}

			v[index] = m11
		case mPickReturnedCards:
			var m12 pickReturnedCardsMove

			if err := json.Unmarshal(*message, &m12); err != nil {
				panic("moves unmarshal fail")
			}

			v[index] = m12
		case mOver:
			var m13 overMove

			if err := json.Unmarshal(*message, &m13); err != nil {
				panic("moves unmarshal fail")
			}

			v[index] = m13
		default:
			panic("unknown move")
		}
	}

	*dst = v

	return nil
}

func UnmarshalMove(move []byte) (Mutator, error) {
	var err error

	var m map[string]interface{}

	if err = json.Unmarshal(move, &m); err != nil {
		return nil, err
	}

	switch moveId(m["id"].(float64)) {
	case mPrepare:
		var m1 prepareMove
		err = json.Unmarshal(move, &m1)

		return m1, err
	case mPickWonder:
		var m2 pickWonderMove
		err = json.Unmarshal(move, &m2)

		return m2, err
	case mPickBoardToken:
		var m3 pickBoardTokenMove
		err = json.Unmarshal(move, &m3)

		return m3, err
	case mConstructCard:
		var m4 constructCardMove
		err = json.Unmarshal(move, &m4)

		return m4, err
	case mConstructWonder:
		var m5 constructWonderMove
		err = json.Unmarshal(move, &m5)

		return m5, err
	case mDiscardCard:
		var m6 discardCardMove
		err = json.Unmarshal(move, &m6)

		return m6, err
	case mSelectWhoBeginsTheNextAge:
		var m7 selectWhoBeginsTheNextAgeMove
		err = json.Unmarshal(move, &m7)

		return m7, err
	case mBurnCard:
		var m8 burnCardMove
		err = json.Unmarshal(move, &m8)

		return m8, err
	case mPickRandomToken:
		var m9 pickRandomTokenMove
		err = json.Unmarshal(move, &m9)

		return m9, err
	case mPickTopLineCard:
		var m10 pickTopLineCardMove
		err = json.Unmarshal(move, &m10)

		return m10, err
	case mPickDiscardedCard:
		var m11 pickDiscardedCardMove
		err = json.Unmarshal(move, &m11)

		return m11, err
	case mPickReturnedCards:
		var m12 pickReturnedCardsMove
		err = json.Unmarshal(move, &m12)

		return m12, err
	case mOver:
		var m13 overMove
		err = json.Unmarshal(move, &m13)

		return m13, err
	default:
		return nil, errors.New("unknown move")
	}
}
