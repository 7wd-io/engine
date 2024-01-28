package engine

const (
	effectBurnCard effectId = iota + 1
	effectChain
	effectCoins
	effectCoinsFor
	effectDiscounter
	effectFine
	effectFixedCost
	effectGuild
	effectMathematics
	effectMilitary
	effectPickDiscardedCard
	effectPickTopLineCard
	effectPickReturnedCards
	effectBoardToken
	effectRandomToken
	effectScience
	effectPoints
	effectResource
	effectPlayAgain
	effectDiscardRewardAdjuster
)

type effectId int

type effect struct {
	Id effectId `json:"id"`
}

func newEffectBurnCard(g cardGroup) burnCardEffect {
	return burnCardEffect{
		effect: effect{effectBurnCard},
		Group:  g,
	}
}

type burnCardEffect struct {
	effect
	Group cardGroup `json:"group"`
}

func (dst burnCardEffect) Mutate(s *State) error {
	items := s.Enemy.Cards.Data[dst.Group]

	if len(items) == 0 {
		// skip effect
		return nil
	}

	s.pushDialog(func(name Nickname) mutator {
		return func(s *State) {
			s.Phase = phaseBurnCard
			s.setTurn(name)

			s.DialogItems.Cards = make(CardList, len(items))
			copy(s.DialogItems.Cards, items)
		}
	}(s.Me.Name))

	return nil
}

func newEffectChain(cid CardId) chainEffect {
	return chainEffect{
		effect: effect{effectChain},
		Card:   cid,
	}
}

type chainEffect struct {
	effect
	Card CardId `json:"card"`
}

func (dst chainEffect) Mutate(s *State) error {
	s.Me.Chains.add(dst.Card)

	return nil
}

func newEffectCoins(count int) coinsEffect {
	return coinsEffect{
		effect: effect{effectCoins},
		Count:  count,
	}
}

type coinsEffect struct {
	effect
	Count int `json:"count"`
}

func (dst coinsEffect) Mutate(s *State) error {
	s.Me.Treasure.add(dst.Count)

	return nil
}

func newEffectCoinsFor(b Bonus, count int) coinsForEffect {
	return coinsForEffect{
		effect: effect{effectCoinsFor},
		Bonus:  b,
		Count:  count,
	}
}

type coinsForEffect struct {
	effect
	Bonus Bonus `json:"bonus"`
	Count int   `json:"count"`
}

func (dst coinsForEffect) Mutate(s *State) error {
	s.Me.Treasure.add(s.Me.bonusRate(dst.Bonus) * dst.Count)

	return nil
}

func newEffectDiscounter(d discountContext, count int, r ...resource) discounterEffect {
	return discounterEffect{
		effect:    effect{effectDiscounter},
		Resources: r,
		count:     count,
		discount:  d,
	}
}

type discounterEffect struct {
	effect
	Resources resourceList `json:"resources"`
	count     int
	discount  discountContext
}

func (dst discounterEffect) Mutate(s *State) error {
	set := make(resourceSet, len(dst.Resources))

	for _, r := range dst.Resources {
		set[r] = struct{}{}
	}

	s.Me.Discounter.add(discount{
		Context:   dst.discount,
		Resources: set,
		Count:     dst.count,
	})

	return nil
}

func newEffectFine(coins int) fineEffect {
	return fineEffect{
		effect: effect{effectFine},
		Coins:  coins,
	}
}

type fineEffect struct {
	effect
	Coins int `json:"coins"`
}

func (dst fineEffect) Mutate(s *State) error {
	s.Enemy.Treasure.add(-dst.Coins)

	return nil
}

func newEffectFixedCost(r ...resource) fixedCostEffect {
	return fixedCostEffect{
		effect:    effect{effectFixedCost},
		Resources: r,
	}
}

type fixedCostEffect struct {
	effect
	Resources resourceList `json:"resources"`
}

func (dst fixedCostEffect) Mutate(s *State) error {
	for _, r := range dst.Resources {
		s.Me.Bank.setResourcePrice(r, fixedResourceCost)
	}

	return nil
}

func newEffectGuild(b Bonus, points int, coins int) guildEffect {
	return guildEffect{
		effect: effect{effectGuild},
		Bonus:  b,
		Points: points,
		Coins:  coins,
	}
}

type guildEffect struct {
	effect
	Bonus  Bonus `json:"bonus"`
	Points int   `json:"points"`
	Coins  int   `json:"coins"`
}

func (dst guildEffect) Mutate(s *State) error {
	s.Me.Treasure.add(dst.rate(s.Me, s.Enemy) * dst.Coins)

	return nil
}

func (dst guildEffect) GetPoints(s *State) int {
	return dst.rate(s.Me, s.Enemy) * dst.Points
}

func (dst guildEffect) rate(c1, c2 *city) int {
	r1 := c1.bonusRate(dst.Bonus)
	r2 := c2.bonusRate(dst.Bonus)

	if r1 >= r2 {
		return r1
	}

	return r2
}

func newEffectMathematics() mathematicsEffect {
	return mathematicsEffect{
		effect: effect{effectMathematics},
	}
}

type mathematicsEffect struct {
	effect
}

func (dst mathematicsEffect) GetPoints(s *State) int {
	return len(s.Me.Tokens.List) * 3
}

type militaryEffectConfig struct {
	power            int
	strategyDisabled bool
}

func newEffectMilitaryWithConfig(c militaryEffectConfig) militaryEffect {
	return militaryEffect{
		effect:           effect{effectMilitary},
		Power:            c.power,
		strategyDisabled: c.strategyDisabled,
	}
}

func newEffectMilitary(power int) militaryEffect {
	return newEffectMilitaryWithConfig(militaryEffectConfig{
		power:            power,
		strategyDisabled: false,
	})
}

type militaryEffect struct {
	effect
	Power            int `json:"power"`
	strategyDisabled bool
}

func (dst militaryEffect) Mutate(s *State) error {
	power := dst.Power

	if !dst.strategyDisabled && s.Me.Tokens.has(Strategy) {
		power++
	}

	fine, supremacy := s.Me.Track.moveConflictPawn(power, s.Enemy.Track)

	if fine > 0 {
		s.Enemy.Treasure.add(-fine)
	}

	if supremacy {
		s.over(MilitarySupremacy, s.Me.Name)
	}

	return nil
}

func newEffectPickDiscardedCard() pickDiscardedCardEffect {
	return pickDiscardedCardEffect{
		effect: effect{effectPickDiscardedCard},
	}
}

type pickDiscardedCardEffect struct {
	effect
}

func (dst pickDiscardedCardEffect) Mutate(s *State) error {
	items := s.CardItems.Discarded

	if len(items) == 0 {
		// skip
		return nil
	}

	s.pushDialog(func(name Nickname) mutator {
		return func(s *State) {
			s.Phase = phasePickDiscardedCard
			s.setTurn(name)
			s.DialogItems.Cards = make(CardList, len(items))
			copy(s.DialogItems.Cards, items)
		}
	}(s.Me.Name))

	return nil
}

func newEffectPickTopLineCard() pickTopLineCardEffect {
	return pickTopLineCardEffect{
		effect: effect{effectPickTopLineCard},
	}
}

type pickTopLineCardEffect struct {
	effect
}

func (dst pickTopLineCardEffect) Mutate(s *State) error {
	items := s.Deck.topLineCards()

	if len(items) == 0 {
		// skip
		return nil
	}

	s.pushDialog(func(name Nickname) mutator {
		return func(s *State) {
			s.Phase = phasePickTopLineCard
			s.setTurn(name)
			s.DialogItems.Cards = items
		}
	}(s.Me.Name))

	return nil
}

func newEffectPickReturnedCards() pickReturnedCardsEffect {
	return pickReturnedCardsEffect{
		effect: effect{effectPickReturnedCards},
	}
}

type pickReturnedCardsEffect struct {
	effect
}

func (dst pickReturnedCardsEffect) Mutate(s *State) error {
	s.pushDialog(func(name Nickname) mutator {
		return func(s *State) {
			s.Phase = phasePickReturnedCards
			s.setTurn(name)
			s.DialogItems.Cards = s.Deck.returnedCards()
		}
	}(s.Me.Name))

	return nil
}

func newEffectPickBoardToken() pickBoardTokenEffect {
	return pickBoardTokenEffect{
		effect: effect{effectBoardToken},
	}
}

type pickBoardTokenEffect struct {
	effect
}

func (dst pickBoardTokenEffect) Mutate(s *State) error {
	s.pushDialog(func(name Nickname) mutator {
		return func(s *State) {
			s.Phase = phasePickBoardToken
			s.setTurn(name)
			s.DialogItems.Tokens = make(TokenList, len(s.Tokens))
			copy(s.DialogItems.Tokens, s.Tokens)
		}
	}(s.Me.Name))

	return nil
}

func newEffectPickRandomToken() pickRandomTokenEffect {
	return pickRandomTokenEffect{
		effect: effect{effectRandomToken},
	}
}

type pickRandomTokenEffect struct {
	effect
}

func (dst pickRandomTokenEffect) Mutate(s *State) error {
	s.pushDialog(func(name Nickname) mutator {
		return func(s *State) {
			s.Phase = phasePickRandomToken
			s.setTurn(name)
			s.DialogItems.Tokens = make(TokenList, len(s.RandomItems.Tokens))
			copy(s.DialogItems.Tokens, s.RandomItems.Tokens)
		}
	}(s.Me.Name))

	return nil
}

func newEffectPlayAgain() playAgainEffect {
	return playAgainEffect{
		effect: effect{effectPlayAgain},
	}
}

type playAgainEffect struct {
	effect
}

func (dst playAgainEffect) Mutate(s *State) error {
	s.PlayAgain = true

	return nil
}

func newEffectPoints(count int) pointsEffect {
	return pointsEffect{
		effect: effect{effectPoints},
		Count:  count,
	}
}

type pointsEffect struct {
	effect
	Count int `json:"count"`
}

func (dst pointsEffect) GetPoints(s *State) int {
	return dst.Count
}

func newEffectResource(r resource, count int) resourceEffect {
	return resourceEffect{
		effect: effect{effectResource},
		Resources: resourceMap{
			r: count,
		},
	}
}

type resourceEffect struct {
	effect
	Resources resourceMap `json:"resources"`
}

func (dst resourceEffect) Mutate(s *State) error {
	for r, count := range dst.Resources {
		s.Me.Resources[r] += count

		if !s.Enemy.Bank.hasFixedResourcePrice(r) {
			s.Enemy.Bank.setResourcePrice(r, defaultResourceCost+s.Me.Resources[r])
		}
	}

	return nil
}

func (dst resourceEffect) Burn(s *State) {
	for r, count := range dst.Resources {
		s.Enemy.Resources[r] -= count

		if !s.Me.Bank.hasFixedResourcePrice(r) {
			s.Me.Bank.setResourcePrice(r, defaultResourceCost+s.Enemy.Resources[r])
		}
	}
}

func newEffectScience(s symbol) scienceEffect {
	return scienceEffect{
		effect: effect{effectScience},
		Symbol: s,
	}
}

type scienceEffect struct {
	effect
	Symbol symbol `json:"symbol"`
}

func (dst scienceEffect) Mutate(s *State) error {
	switch s.Me.Symbols.add(dst.Symbol) {
	case symbolStatusToken:
		return newEffectPickBoardToken().Mutate(s)
	case symbolStatusSupremacy:
		s.over(ScienceSupremacy, s.Me.Name)
	}

	return nil
}

func newEffectDiscardRewardAdjuster() discardRewardAdjusterEffect {
	return discardRewardAdjusterEffect{
		effect: effect{effectDiscardRewardAdjuster},
	}
}

type discardRewardAdjusterEffect struct {
	effect
}

func (dst discardRewardAdjusterEffect) Mutate(s *State) error {
	s.Me.Bank.setDiscardReward(s.Me.Bank.getDiscardReward() + 1)

	return nil
}
