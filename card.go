package engine

const (
	LumberYard CardId = iota + 100
	LoggingCamp
	ClayPool
	ClayPit
	Quarry
	StonePit
	GlassWorks
	Press
	GuardTower
	Workshop
	Apothecary
	StoneReserve
	ClayReserve
	WoodReserve
	Stable
	Garrison
	Palisade
	Scriptorium
	Pharmacist
	Theater
	Altar
	Baths
	Tavern
)

const (
	SawMill CardId = iota + 200
	BrickYard
	ShelfQuarry
	GlassBlower
	DryingRoom
	Walls
	Forum
	Caravansery
	CustomHouse
	CourtHouse
	HorseBreeders
	Barracks
	ArcheryRange
	ParadeGround
	Library
	Dispensary
	School
	Laboratory
	Statue
	Temple
	Aqueduct
	Rostrum
	Brewery
)

const (
	Arsenal CardId = iota + 300
	Pretorium
	Academy
	Study
	ChamberOfCommerce
	Port
	Armory
	Palace
	TownHall
	Obelisk
	Fortifications
	SiegeWorkshop
	Circus
	University
	Observatory
	Gardens
	Pantheon
	Senate
	Lighthouse
	Arena
)

const (
	MerchantsGuild CardId = iota + 400
	ShipOwnersGuild
	BuildersGuild
	MagistratesGuild
	ScientistsGuild
	MoneyLendersGuild
	TacticiansGuild
)

const (
	groupRawMaterials cardGroup = iota + 1
	groupManufacturedGoods
	groupMilitary
	groupScientific
	groupCivilian
	groupCommercial
	groupGuild
)

type CardId int

func (dst CardId) isNil() bool {
	return dst == 0
}

type CardMap map[CardId]Card
type CardList []CardId
type CardSet map[CardId]struct{}

func (dst CardSet) List() CardList {
	out := make(CardList, len(dst))

	i := 0

	for cid, _ := range dst {
		out[i] = cid
		i++
	}

	return out
}

type cardGroup int

type Card struct {
	Id    CardId    `json:"id"`
	Age   age       `json:"age"`
	Group cardGroup `json:"group"`
	Cost  cost      `json:"cost,omitempty"`
	unit
}
