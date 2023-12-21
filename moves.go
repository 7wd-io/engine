package engine

import (
	"encoding/json"
	"errors"
	"log"
)

const (
	MovePrepare MoveId = iota + 1
	MovePickWonder
	MovePickBoardToken
	MoveConstructCard
	MoveConstructWonder
	MoveDiscardCard
	MoveSelectWhoBeginsTheNextAge
	MoveBurnCard
	MovePickRandomToken
	MovePickTopLineCard
	MovePickDiscardedCard
	MovePickReturnedCards
	MoveOver
)

type MoveId int

type move struct {
	Id MoveId `json:"id"`
}

func NewMovePrepare(p1, p2 Nickname) Mutator {
	if randBool() {
		p2, p1 = p1, p2
	}

	t := shuffleTokens(R.tokenIds)

	w := shuffleWonders(R.wonderIds)

	c := map[age]CardList{}

	for _, a := range []age{ageI, ageII, ageIII} {
		c[a] = randCards(a)
	}

	return PrepareMove{
		Id:           MovePrepare,
		P1:           p1,
		P2:           p2,
		Wonders:      w[:wonderSelectionPoolSize*2],
		Tokens:       t[:startingTokensCount],
		RandomTokens: t[startingTokensCount : startingTokensCount+randomTokensCount],
		Cards:        c,
	}
}

type PrepareMove struct {
	Id           MoveId           `json:"id"`
	P1           Nickname         `json:"p1"`
	P2           Nickname         `json:"p2"`
	Wonders      WonderList       `json:"wonders"`
	Tokens       TokenList        `json:"tokens"`
	RandomTokens TokenList        `json:"randomTokens"`
	Cards        map[age]CardList `json:"cards"`
}

func (dst PrepareMove) Mutate(s *State) error {
	if s.Phase != phaseNil {
		return ErrActionNotAllowed
	}

	s.Age = ageI
	s.Phase = phasePrepare
	s.FirstTurn = dst.P1

	s.Tokens = make(TokenList, len(dst.Tokens))
	copy(s.Tokens, dst.Tokens)

	s.Me = newCity(dst.P1)
	s.Enemy = newCity(dst.P2)

	s.Dialogs = []mutator{}
	s.RandomItems.Cards = dst.Cards
	s.RandomItems.Tokens = dst.RandomTokens
	s.RandomItems.Wonders = dst.Wonders
	s.DialogItems.Wonders = make(WonderList, wonderSelectionPoolSize)
	copy(s.DialogItems.Wonders, dst.Wonders[:wonderSelectionPoolSize])

	return nil
}

func NewMovePickWonder(w WonderId) Mutator {
	return PickWonderMove{
		Id:     MovePickWonder,
		Wonder: w,
	}
}

type PickWonderMove struct {
	Id     MoveId   `json:"id"`
	Wonder WonderId `json:"wonder"`
}

func (dst PickWonderMove) Mutate(s *State) error {
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
	return SelectWhoBeginsTheNextAgeMove{
		Id:     MoveSelectWhoBeginsTheNextAge,
		Player: n,
	}
}

type SelectWhoBeginsTheNextAgeMove struct {
	Id     MoveId   `json:"id"`
	Player Nickname `json:"player"`
}

func (dst SelectWhoBeginsTheNextAgeMove) Mutate(s *State) error {
	if s.Phase != phaseSelectWhoBeginsTheNextAge {
		return ErrActionNotAllowed
	}

	s.setTurn(dst.Player)

	s.Phase = PhaseTurn

	return nil
}

func NewMovePickBoardToken(t TokenId) Mutator {
	return PickBoardTokenMove{
		Id:    MovePickBoardToken,
		Token: t,
	}
}

type PickBoardTokenMove struct {
	Id    MoveId  `json:"id"`
	Token TokenId `json:"token"`
}

func (dst PickBoardTokenMove) Mutate(s *State) error {
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

	if err := R.tokens[dst.Token].Mutate(s); err != nil {
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
	return PickRandomTokenMove{
		Id:    MovePickRandomToken,
		Token: t,
	}
}

type PickRandomTokenMove struct {
	Id    MoveId  `json:"id"`
	Token TokenId `json:"token"`
}

func (dst PickRandomTokenMove) Mutate(s *State) error {
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

	if err := R.tokens[dst.Token].Mutate(s); err != nil {
		return err
	}

	s.after()

	return nil
}

func NewMoveConstructCard(c CardId) Mutator {
	return ConstructCardMove{
		Id:   MoveConstructCard,
		Card: c,
	}
}

type ConstructCardMove struct {
	Id   MoveId `json:"id"`
	Card CardId `json:"card"`
}

func (dst ConstructCardMove) Mutate(s *State) error {
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

	if err := R.Cards[dst.Card].Mutate(s); err != nil {
		return err
	}

	s.after()

	return nil
}

func NewMoveConstructWonder(w WonderId, c CardId) Mutator {
	return ConstructWonderMove{
		Id:     MoveConstructWonder,
		Wonder: w,
		Card:   c,
	}
}

type ConstructWonderMove struct {
	Id     MoveId   `json:"id"`
	Wonder WonderId `json:"wonder"`
	Card   CardId   `json:"card"`
}

func (dst ConstructWonderMove) Mutate(s *State) error {
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

	if err := R.Wonders[dst.Wonder].Mutate(s); err != nil {
		return err
	}

	if s.Me.Tokens.has(Theology) {
		s.PlayAgain = true
	}

	s.after()

	return nil
}

func NewMoveDiscardCard(c CardId) Mutator {
	return DiscardCardMove{
		Id:   MoveDiscardCard,
		Card: c,
	}
}

type DiscardCardMove struct {
	Id   MoveId `json:"id"`
	Card CardId `json:"card"`
}

func (dst DiscardCardMove) Mutate(s *State) error {
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
	return BurnCardMove{
		Id:   MoveBurnCard,
		Card: c,
	}
}

type BurnCardMove struct {
	Id   MoveId `json:"id"`
	Card CardId `json:"card"`
}

func (dst BurnCardMove) Mutate(s *State) error {
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

	R.Cards[dst.Card].Burn(s)

	s.after()

	return nil
}

func NewMovePickTopLineCard(c CardId) Mutator {
	return PickTopLineCardMove{
		Id:   MovePickTopLineCard,
		Card: c,
	}
}

type PickTopLineCardMove struct {
	Id   MoveId `json:"id"`
	Card CardId `json:"card"`
}

func (dst PickTopLineCardMove) Mutate(s *State) error {
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

	if err := R.Cards[dst.Card].Mutate(s); err != nil {
		return err
	}

	s.after()

	return nil
}

func NewMovePickDiscardedCard(c CardId) Mutator {
	return PickDiscardedCardMove{
		Id:   MovePickDiscardedCard,
		Card: c,
	}
}

type PickDiscardedCardMove struct {
	Id   MoveId `json:"id"`
	Card CardId `json:"card"`
}

func (dst PickDiscardedCardMove) Mutate(s *State) error {
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

	if err := R.Cards[dst.Card].Mutate(s); err != nil {
		return err
	}

	s.after()

	return nil
}

func NewMovePickReturnedCards(p CardId, g CardId) Mutator {
	return PickReturnedCardsMove{
		Id:   MovePickReturnedCards,
		Pick: p,
		Give: g,
	}
}

type PickReturnedCardsMove struct {
	Id   MoveId `json:"id"`
	Pick CardId `json:"pick"`
	Give CardId `json:"give"`
}

func (dst PickReturnedCardsMove) Mutate(s *State) error {
	if s.Phase != phasePickReturnedCards {
		return ErrActionNotAllowed
	}

	if !dst.isValidCards(s) {
		return ErrActionNotAllowed
	}

	s.byCity(s.Me.Name, func() {
		s.Me.Cards.add(dst.Pick)
		_ = R.Cards[dst.Pick].Mutate(s)
	})

	s.byCity(s.Enemy.Name, func() {
		s.Me.Cards.add(dst.Give)
		_ = R.Cards[dst.Give].Mutate(s)
	})

	s.after()

	return nil
}

func (dst PickReturnedCardsMove) isValidCards(s *State) bool {
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
	return OverMove{
		Id:     MoveOver,
		Loser:  loser,
		Reason: reason,
	}
}

type OverMove struct {
	Id     MoveId   `json:"id"`
	Loser  Nickname `json:"loser"`
	Reason Victory  `json:"reason"`
}

func (dst OverMove) Mutate(s *State) error {
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

		switch MoveId(m["id"].(float64)) {
		case MovePrepare:
			var m1 PrepareMove

			if err := json.Unmarshal(*message, &m1); err != nil {
				panic("moves unmarshal fail")
			}

			v[index] = m1
		case MovePickWonder:
			var m2 PickWonderMove

			if err := json.Unmarshal(*message, &m2); err != nil {
				panic("moves unmarshal fail")
			}

			v[index] = m2
		case MovePickBoardToken:
			var m3 PickBoardTokenMove

			if err := json.Unmarshal(*message, &m3); err != nil {
				panic("moves unmarshal fail")
			}

			v[index] = m3
		case MoveConstructCard:
			var m4 ConstructCardMove

			if err := json.Unmarshal(*message, &m4); err != nil {
				panic("moves unmarshal fail")
			}

			v[index] = m4
		case MoveConstructWonder:
			var m5 ConstructWonderMove

			if err := json.Unmarshal(*message, &m5); err != nil {
				panic("moves unmarshal fail")
			}

			v[index] = m5
		case MoveDiscardCard:
			var m6 DiscardCardMove

			if err := json.Unmarshal(*message, &m6); err != nil {
				panic("moves unmarshal fail")
			}

			v[index] = m6
		case MoveSelectWhoBeginsTheNextAge:
			var m7 SelectWhoBeginsTheNextAgeMove

			if err := json.Unmarshal(*message, &m7); err != nil {
				panic("moves unmarshal fail")
			}

			v[index] = m7
		case MoveBurnCard:
			var m8 BurnCardMove

			if err := json.Unmarshal(*message, &m8); err != nil {
				panic("moves unmarshal fail")
			}

			v[index] = m8
		case MovePickRandomToken:
			var m9 PickRandomTokenMove

			if err := json.Unmarshal(*message, &m9); err != nil {
				panic("moves unmarshal fail")
			}

			v[index] = m9
		case MovePickTopLineCard:
			var m10 PickTopLineCardMove

			if err := json.Unmarshal(*message, &m10); err != nil {
				panic("moves unmarshal fail")
			}

			v[index] = m10
		case MovePickDiscardedCard:
			var m11 PickDiscardedCardMove

			if err := json.Unmarshal(*message, &m11); err != nil {
				panic("moves unmarshal fail")
			}

			v[index] = m11
		case MovePickReturnedCards:
			var m12 PickReturnedCardsMove

			if err := json.Unmarshal(*message, &m12); err != nil {
				panic("moves unmarshal fail")
			}

			v[index] = m12
		case MoveOver:
			var m13 OverMove

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

	switch MoveId(m["id"].(float64)) {
	case MovePrepare:
		var m1 PrepareMove
		err = json.Unmarshal(move, &m1)

		return m1, err
	case MovePickWonder:
		var m2 PickWonderMove
		err = json.Unmarshal(move, &m2)

		return m2, err
	case MovePickBoardToken:
		var m3 PickBoardTokenMove
		err = json.Unmarshal(move, &m3)

		return m3, err
	case MoveConstructCard:
		var m4 ConstructCardMove
		err = json.Unmarshal(move, &m4)

		return m4, err
	case MoveConstructWonder:
		var m5 ConstructWonderMove
		err = json.Unmarshal(move, &m5)

		return m5, err
	case MoveDiscardCard:
		var m6 DiscardCardMove
		err = json.Unmarshal(move, &m6)

		return m6, err
	case MoveSelectWhoBeginsTheNextAge:
		var m7 SelectWhoBeginsTheNextAgeMove
		err = json.Unmarshal(move, &m7)

		return m7, err
	case MoveBurnCard:
		var m8 BurnCardMove
		err = json.Unmarshal(move, &m8)

		return m8, err
	case MovePickRandomToken:
		var m9 PickRandomTokenMove
		err = json.Unmarshal(move, &m9)

		return m9, err
	case MovePickTopLineCard:
		var m10 PickTopLineCardMove
		err = json.Unmarshal(move, &m10)

		return m10, err
	case MovePickDiscardedCard:
		var m11 PickDiscardedCardMove
		err = json.Unmarshal(move, &m11)

		return m11, err
	case MovePickReturnedCards:
		var m12 PickReturnedCardsMove
		err = json.Unmarshal(move, &m12)

		return m12, err
	case MoveOver:
		var m13 OverMove
		err = json.Unmarshal(move, &m13)

		return m13, err
	default:
		return nil, errors.New("unknown move")
	}
}
