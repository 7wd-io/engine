package engine

const (
	clay resource = iota + 1
	wood
	stone
	glass
	papyrus
)

type resource int
type resourceList []resource
type resourceSet map[resource]struct{}
type resourceMap map[resource]int

type cost struct {
	Coins     int         `json:"coins"`
	Resources resourceMap `json:"resources"`
}

func (dst *cost) makeCopy() cost {
	rmap := make(resourceMap, len(dst.Resources))

	for r, count := range dst.Resources {
		rmap[r] = count
	}

	return cost{
		Coins:     dst.Coins,
		Resources: rmap,
	}
}

func (dst *cost) sub(sub resourceMap) {
	for r, count := range dst.Resources {
		if value, ok := sub[r]; ok {
			dst.Resources[r] = max(count-value, 0)
		}
	}
}
