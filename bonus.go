package engine

const (
	BonusResources Bonus = iota + 1
	BonusRawMaterials
	BonusManufacturedGoods
	BonusMilitary
	BonusCommercial
	BonusCivilian
	BonusScience
	BonusWonder
	BonusCoin
)

type Bonus int
