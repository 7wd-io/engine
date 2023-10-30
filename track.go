package main

const capitalPos = 9

const (
	zone1 trackZoneInd = iota
	zone2
	zone3
	zone4
)

var trackZones = map[trackZoneInd]trackZone{
	zone1: {
		startPos: 0,
		points:   0,
		fine:     0,
	},
	zone2: {
		startPos: 1,
		points:   2,
		fine:     0,
	},
	zone3: {
		startPos: 3,
		points:   5,
		fine:     2,
	},
	zone4: {
		startPos: 6,
		points:   10,
		fine:     5,
	},
}

type track struct {
	Pos     int          `json:"pos"`
	MaxZone trackZoneInd `json:"maxZone"`
}

func (dst *track) moveConflictPawn(power int, enemy *track) (fine int, supremacy bool) {
	if enemy.Pos >= power {
		enemy.Pos -= power
		return fine, supremacy
	}

	dst.Pos += power - enemy.Pos
	enemy.Pos = 0

	if dst.Pos >= capitalPos {
		dst.Pos = capitalPos
		supremacy = true
	}

	if ind := getTrackZoneInd(dst.Pos); ind > dst.MaxZone {
		dst.MaxZone = ind
		fine = trackZones[ind].fine
	}

	return fine, supremacy
}

func (dst *track) getPoints() int {
	return trackZones[getTrackZoneInd(dst.Pos)].points
}

type trackZoneInd int

type trackZone struct {
	startPos int
	points   int
	fine     int
}

func getTrackZoneInd(pos int) trackZoneInd {
	for _, ind := range []trackZoneInd{zone4, zone3, zone2} {
		z := trackZones[ind]

		if pos >= z.startPos {
			return ind
		}
	}

	return zone1
}
