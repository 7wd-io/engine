package main

const (
	dcGlobal discountContext = iota + 1
	dcCivilian
	dcWonders
)

type discountContext int

type discount struct {
	Context   discountContext `json:"context"`
	Resources resourceSet     `json:"resources"`
	Count     int             `json:"-"`
}

func (dst *discount) isSupport(ctx discountContext) bool {
	if dst.Context == dcGlobal {
		return true
	}

	return dst.Context == ctx
}

func (dst *discount) discount(cost resourceMap, priority resourceList) {
	var dsc int
	reserve := dst.Count

	for _, r := range priority {
		_, has := dst.Resources[r]

		if !has {
			continue
		}

		if reserve == 0 {
			break
		}

		if cost[r] > 0 {
			if cost[r] < reserve {
				dsc = cost[r]
			} else {
				dsc = reserve
			}

			cost[r] -= dsc
			reserve -= dsc
		}
	}
}

func getDiscountContextByCard(cid CardId) discountContext {
	if R.Cards[cid].Group == groupCivilian {
		return dcCivilian
	}

	return dcGlobal
}
