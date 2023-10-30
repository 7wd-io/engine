package main

const (
	TheAppianWay WonderId = iota + 1
	CircusMaximus
	TheColossus
	TheGreatLibrary
	TheGreatLighthouse
	TheHangingGardens
	TheMausoleum
	Piraeus
	ThePyramids
	TheSphinx
	TheStatueOfZeus
	TheTempleOfArtemis
	Messe
	StatueOfLiberty
)

type WonderId int
type WonderMap map[WonderId]wonder
type wonderSearch map[WonderId]struct{}
type wonderList []WonderId

func (dst wonderList) getWonderSearch() wonderSearch {
	out := make(wonderSearch, len(dst))

	for _, wid := range dst {
		out[wid] = struct{}{}
	}

	return out
}

type wonder struct {
	Id   WonderId `json:"id"`
	Cost cost     `json:"cost"`
	unit
}
