package engine

import (
	"sort"
)

type State struct {
	Age       age       `json:"age"`
	Phase     phase     `json:"phase"`
	FirstTurn Nickname  `json:"firstTurn"`
	Tokens    tokenList `json:"tokens"`
	Me        *city     `json:"me"`
	Enemy     *city     `json:"enemy"`

	Dialogs      []mutator `json:"-"`
	PlayAgain    bool      `json:"-"`
	FallbackTurn Nickname  `json:"-"`

	CardItems   cardItems   `json:"cardItems"`
	DialogItems dialogItems `json:"dialogItems"`

	Winner  *Nickname `json:"winner,omitempty"`
	Victory *Victory  `json:"victory,omitempty"`

	RandomItems randomItems `json:"-"`
	Deck        *deck       `json:"-"`
}

func (dst *State) pushDialog(d mutator) {
	dst.Dialogs = append(dst.Dialogs, d)
}

func (dst *State) pullDialog() mutator {
	var d mutator
	d, dst.Dialogs = dst.Dialogs[0], dst.Dialogs[1:]

	return d
}

func (dst *State) byCity(nickname Nickname, fn func()) {
	orig := dst.Me.Name

	dst.setTurn(nickname)

	fn()

	dst.setTurn(orig)
}

func (dst *State) over(v Victory, winner Nickname) {
	// we can handle two units for both players at the same time (Statue of Liberty)
	// and each unit can call over
	// whoever called first wins
	if dst.Phase == phaseOver {
		return
	}

	dst.Phase = phaseOver

	dst.refreshCities()

	if winner == "" {
		if dst.Me.Score.Total != dst.Enemy.Score.Total {
			if dst.Me.Score.Total > dst.Enemy.Score.Total {
				winner = dst.Me.Name
			} else {
				winner = dst.Enemy.Name
			}
		} else if dst.Me.Score.Civilian != dst.Enemy.Score.Civilian {
			if dst.Me.Score.Civilian > dst.Enemy.Score.Civilian {
				winner = dst.Me.Name
			} else {
				winner = dst.Enemy.Name
			}
		} else {
			if dst.Me.Name == dst.FirstTurn {
				winner = dst.Enemy.Name
			} else {
				winner = dst.Me.Name
			}
		}
	}

	dst.Winner = &winner
	dst.Victory = &v
}

func (dst *State) pay(d discountContext, cost cost) error {
	price := dst.Me.finalPrice(d, cost)

	if price > dst.Me.Treasure.Coins {
		return ErrActionNotAllowed
	}

	dst.Me.Treasure.add(-price)

	if dst.Enemy.Tokens.has(Economy) {
		dst.Enemy.Treasure.add(price - cost.Coins)
	}

	return nil
}

func (dst *State) nextTurn() {
	dst.Me, dst.Enemy = dst.Enemy, dst.Me
}

func (dst *State) setTurn(n Nickname) {
	if dst.Enemy.Name == n {
		dst.nextTurn()
	}
}

func (dst *State) enemyFor(p Nickname) Nickname {
	if dst.Me.Name == p {
		return dst.Enemy.Name
	}

	return dst.Me.Name
}

func (dst *State) after() {
	if dst.IsOver() {
		return
	}

	var isOver bool

	hasDialogs := len(dst.Dialogs) > 0

	isOver = dst.Deck.isEmpty() && dst.Age == ageIII && dst.Phase == PhaseTurn && !hasDialogs

	if isOver {
		dst.over(Civilian, "")
		return
	}

	// dialog has own logic to set turn
	// we resolve turn and use this to fallback turn after dialog if needed
	dst.resolveTurn()

	if !hasDialogs {
		dst.nextAgeIfNeeded()
	}

	dst.refreshDeckCards()
	dst.refreshCities()

	if hasDialogs {
		if dst.FallbackTurn == "" {
			dst.FallbackTurn = dst.Me.Name
		}

		dialog := dst.pullDialog()
		dialog(dst)

		return
	} else if dst.FallbackTurn != "" {
		// if starts next age, origin turn resolve is priority
		if dst.Phase != phaseSelectWhoBeginsTheNextAge {
			dst.Phase = PhaseTurn
			dst.setTurn(dst.FallbackTurn)
		}

		dst.FallbackTurn = ""
	}

	isOver = dst.Deck.isEmpty() && dst.Age == ageIII && dst.Phase == PhaseTurn

	if isOver {
		dst.over(Civilian, "")
	}
}

func (dst *State) IsOver() bool {
	return dst.Phase == phaseOver
}

func (dst *State) resolveTurn() {
	if dst.Deck.isEmpty() && dst.Age != ageIII {
		dst.Phase = phaseSelectWhoBeginsTheNextAge
		// if last card construct wonder with play again effect
		// we should drop
		dst.PlayAgain = false

		if dst.Me.Track.Pos == dst.Enemy.Track.Pos {
			return
		}

		if dst.Enemy.Track.Pos == 0 {
			dst.nextTurn()
		}

		// otherwise, me stay

		return
	}

	if dst.PlayAgain {
		dst.PlayAgain = false

		return
	}

	dst.nextTurn()
}

func (dst *State) nextAgeIfNeeded() {
	if dst.Deck.isEmpty() && dst.Age != ageIII {
		dst.Age++
		dst.Deck = newDeck(dst.RandomItems.Cards[dst.Age], R.layouts[dst.Age])
	}
}

func (dst *State) refreshDeckCards() {
	dst.CardItems.Layout = dst.Deck.getLayout()
	dst.CardItems.Playable = dst.Deck.playableCards()
}

func (dst *State) refreshCities() {
	turn := dst.Me.Name

	for _, name := range []Nickname{dst.Me.Name, dst.Enemy.Name} {
		// change current city for right calculates which use Me,Enemy links
		dst.setTurn(name)
		city := dst.Me
		city.refreshCardsPrice(dst.CardItems.Playable)
		city.refreshWondersPrice()
		city.refreshScore(dst)
	}

	// restore original value
	dst.setTurn(turn)
}

func newCity(name Nickname) *city {
	return &city{
		Name:      name,
		Resources: resourceMap{},
		Wonders: &cwonders{
			List:        wonderList{},
			Constructed: map[WonderId]CardId{},
		},
		Tokens: &ctokens{
			List:   tokenList{},
			search: map[TokenId]struct{}{},
		},
		Symbols: &csymbols{
			Data:  map[symbol]int{},
			Order: []symbol{},
		},
		Cards: &ccards{
			Data: map[cardGroup]cardList{},
		},
		Treasure: &ctreasure{
			Coins: defaultTreasureCoins,
		},
		Chains: &cchains{
			List:   cardList{},
			search: map[CardId]struct{}{},
		},
		Track: &track{},
		Bank: bank{
			DiscardReward: defaultDiscardReward,
			ResourcePrice: resourceMap{
				clay:    defaultResourceCost,
				stone:   defaultResourceCost,
				wood:    defaultResourceCost,
				glass:   defaultResourceCost,
				papyrus: defaultResourceCost,
			},
		},
	}
}

type city struct {
	Name       Nickname    `json:"name"`
	Score      score       `json:"score"`
	Resources  resourceMap `json:"resources"`
	Wonders    *cwonders   `json:"wonders"`
	Tokens     *ctokens    `json:"tokens"`
	Symbols    *csymbols   `json:"symbols"`
	Cards      *ccards     `json:"cards"`
	Treasure   *ctreasure  `json:"treasure"`
	Chains     *cchains    `json:"chains"`
	Track      *track      `json:"track"`
	Discounter discounter  `json:"discounter"`
	Bank       bank        `json:"bank"`
}

func (dst *city) refreshCardsPrice(playable cardSet) {
	p := make(map[CardId]int)

	for cid, _ := range playable {
		if dst.Chains.has(cid) {
			p[cid] = 0
			continue
		}

		p[cid] = dst.finalPrice(
			getDiscountContextByCard(cid),
			R.Cards[cid].Cost,
		)
	}

	dst.Bank.CardPrice = p
}

func (dst *city) refreshWondersPrice() {
	p := make(map[WonderId]int)

	for _, wid := range dst.Wonders.List {
		if dst.Wonders.isConstructed(wid) {
			continue
		}

		p[wid] = dst.finalPrice(dcWonders, R.Wonders[wid].Cost)
	}

	dst.Bank.WonderPrice = p
}

func (dst *city) refreshScore(s *State) {
	scr := score{}

	for group, cids := range dst.Cards.Data {
		for _, cid := range cids {
			c := R.Cards[cid]
			points := c.GetPoints(s)

			switch group {
			case groupCivilian:
				scr.Civilian += points
			case groupScientific:
				scr.Science += points
			case groupCommercial:
				scr.Commercial += points
			case groupGuild:
				scr.Guilds += points
			}
		}
	}

	for wid, cid := range dst.Wonders.Constructed {
		if !cid.isNil() {
			scr.Wonders += R.Wonders[wid].GetPoints(s)
		}
	}

	for _, tid := range dst.Tokens.List {
		scr.Tokens += R.tokens[tid].GetPoints(s)
	}

	scr.Coins = dst.Treasure.Coins / coinsPerPoint
	scr.Military = dst.Track.getPoints()
	scr.Total = scr.Civilian +
		scr.Science +
		scr.Commercial +
		scr.Guilds +
		scr.Wonders +
		scr.Tokens +
		scr.Coins +
		scr.Military

	dst.Score = scr
}

func (dst *city) finalPrice(ctx discountContext, c cost) int {
	var price int
	copyCost := c.makeCopy()
	copyCost.sub(dst.Resources)

	dst.Discounter.discount(ctx, copyCost.Resources, dst.Bank.ResourcePrice)

	for r, count := range copyCost.Resources {
		price += dst.Bank.ResourcePrice[r] * count
	}

	return price + copyCost.Coins
}

func (dst *city) bonusRate(b bonus) int {
	switch b {
	case bonusResources:
		return dst.bonusRate(bonusRawMaterials) + dst.bonusRate(bonusManufacturedGoods)
	case bonusRawMaterials:
		return len(dst.Cards.Data[groupRawMaterials])
	case bonusManufacturedGoods:
		return len(dst.Cards.Data[groupManufacturedGoods])
	case bonusMilitary:
		return len(dst.Cards.Data[groupMilitary])
	case bonusCommercial:
		return len(dst.Cards.Data[groupCommercial])
	case bonusCivilian:
		return len(dst.Cards.Data[groupCivilian])
	case bonusScience:
		return len(dst.Cards.Data[groupScientific])
	case bonusWonder:
		var rate int

		for _, c := range dst.Wonders.Constructed {
			if c != 0 {
				rate++
			}
		}

		return rate
	case bonusCoin:
		return dst.Treasure.Coins / coinsPerPoint
	default:
		panic("unsupported bonus")
	}
}

type bank struct {
	DiscardReward int              `json:"discardReward"`
	CardPrice     map[CardId]int   `json:"cardPrice"`
	WonderPrice   map[WonderId]int `json:"wonderPrice"`
	ResourcePrice map[resource]int `json:"resourcePrice"`
}

func (dst *bank) setDiscardReward(coins int) {
	dst.DiscardReward = coins
}

func (dst *bank) getDiscardReward() int {
	return dst.DiscardReward
}

func (dst *bank) setResourcePrice(r resource, coins int) {
	dst.ResourcePrice[r] = coins
}

func (dst *bank) getResourcePrice(r resource) int {
	return dst.ResourcePrice[r]
}

func (dst *bank) hasFixedResourcePrice(r resource) bool {
	return dst.ResourcePrice[r] == fixedResourceCost
}

func (dst *bank) setCardPrice(c CardId, coins int) {
	dst.CardPrice[c] = coins
}

func (dst *bank) setWonderPrice(w WonderId, coins int) {
	dst.WonderPrice[w] = coins
}

type cwonders struct {
	List        wonderList          `json:"list"`
	Constructed map[WonderId]CardId `json:"constructed"`
}

func (dst *cwonders) add(w WonderId) {
	dst.List = append(dst.List, w)
	dst.Constructed[w] = 0
}

func (dst *cwonders) has(w WonderId) bool {
	_, ok := dst.Constructed[w]

	return ok
}

func (dst *cwonders) countTotal() int {
	return len(dst.List)
}

func (dst *cwonders) countConstructed() int {
	var count int

	for _, cid := range dst.Constructed {
		if !cid.isNil() {
			count++
		}
	}

	return count
}

func (dst *cwonders) isConstructed(w WonderId) bool {
	cid, _ := dst.Constructed[w]

	return !cid.isNil()
}

func (dst *cwonders) construct(w WonderId, c CardId) {
	dst.Constructed[w] = c
}

func (dst *cwonders) removeNotConstructed() {
	w := wonderList{}

	for _, wid := range dst.List {
		if dst.Constructed[wid].isNil() {
			delete(dst.Constructed, wid)
		} else {
			w = append(w, wid)
		}
	}

	dst.List = w
}

type ctokens struct {
	List   tokenList `json:"list"`
	search map[TokenId]struct{}
}

func (dst *ctokens) add(t TokenId) {
	dst.List = append(dst.List, t)
	dst.search[t] = struct{}{}
}

func (dst *ctokens) has(t TokenId) bool {
	_, ok := dst.search[t]

	return ok
}

type symbolStatus int

const (
	symbolStatusNil = iota
	symbolStatusToken
	symbolStatusSupremacy
)

type csymbols struct {
	Data  map[symbol]int `json:"data"`
	Order []symbol       `json:"order"`
}

func (dst *csymbols) add(s symbol) symbolStatus {
	if _, ok := dst.Data[s]; !ok {
		dst.Order = append(dst.Order, s)
	}

	dst.Data[s]++

	if dst.Data[s] == symbolCountToken {
		return symbolStatusToken
	}

	if len(dst.Data) == symbolCountSupremacy {
		return symbolStatusSupremacy
	}

	return symbolStatusNil
}

// @TODO make custom map type
type ccards struct {
	Data map[cardGroup]cardList `json:"data"`
}

func (dst *ccards) add(cid CardId) {
	c, _ := R.Cards[cid]

	dst.Data[c.Group] = append(dst.Data[c.Group], cid)
}

func (dst *ccards) remove(cid CardId) {
	c, _ := R.Cards[cid]

	var pos int

	for i, item := range dst.Data[c.Group] {
		if item == cid {
			pos = i
			break
		}
	}

	dst.Data[c.Group] = append(dst.Data[c.Group][:pos], dst.Data[c.Group][pos+1:]...)
}

type ctreasure struct {
	Coins int `json:"coins"`
}

func (dst *ctreasure) add(count int) {
	dst.Coins += count

	if dst.Coins < 0 {
		dst.Coins = 0
	}
}

type cchains struct {
	List   cardList `json:"list"`
	search map[CardId]struct{}
}

func (dst *cchains) add(c CardId) {
	dst.List = append(dst.List, c)
	dst.search[c] = struct{}{}
}

func (dst *cchains) has(c CardId) bool {
	_, ok := dst.search[c]

	return ok
}

type dialogItems struct {
	Wonders wonderList `json:"wonders"`
	Cards   cardList   `json:"cards"`
	Tokens  tokenList  `json:"tokens"`
}

type cardItems struct {
	Layout    cardList `json:"layout"`
	Playable  cardSet  `json:"playable"`
	Discarded cardList `json:"discarded"`
}

func (dst *cardItems) isPlayable(c CardId) bool {
	_, ok := dst.Playable[c]

	return ok
}

func (dst *cardItems) addDiscarded(c CardId) {
	dst.Discarded = append(dst.Discarded, c)
}

func (dst *cardItems) removeDiscarded(c CardId) {
	var pos int

	for i, cid := range dst.Discarded {
		if cid == c {
			pos = i
		}

		break
	}

	dst.Discarded = append(dst.Discarded[:pos], dst.Discarded[pos+1:]...)
}

type randomItems struct {
	Wonders wonderList
	Cards   map[age]cardList
	Tokens  tokenList
}

type discounter struct {
	List []discount `json:"list"`
}

func (dst *discounter) add(d discount) {
	dst.List = append(dst.List, d)
}

func (dst *discounter) discount(ctx discountContext, cost resourceMap, priceList map[resource]int) {
	for _, item := range dst.List {
		if item.isSupport(ctx) {
			item.discount(cost, dst.getPriority(priceList))
		}
	}
}

func (dst *discounter) getPriority(priceList map[resource]int) resourceList {
	var data []struct {
		id    resource
		price int
	}

	for r, price := range priceList {
		data = append(data, struct {
			id    resource
			price int
		}{
			id:    r,
			price: price,
		})
	}

	sort.SliceStable(data, func(i, j int) bool {
		return data[i].price > data[j].price
	})

	sorted := make(resourceList, len(data))

	for _, item := range data {
		sorted = append(sorted, item.id)
	}

	return sorted
}

type score struct {
	Civilian   int `json:"civilian"`
	Science    int `json:"science"`
	Commercial int `json:"commercial"`
	Guilds     int `json:"guilds"`
	Wonders    int `json:"wonders"`
	Tokens     int `json:"tokens"`
	Coins      int `json:"coins"`
	Military   int `json:"military"`
	Total      int `json:"total"`
}
