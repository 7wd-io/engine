package engine

const (
	bonusResources bonus = iota + 1
	bonusRawMaterials
	bonusManufacturedGoods
	bonusMilitary
	bonusCommercial
	bonusCivilian
	bonusScience
	bonusWonder
	bonusCoin
)

type bonus int
