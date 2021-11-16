package recipes

type MaterialCost struct {
	Resources string
	Refined string
	Reagent string
}

type Recipe struct {
	ID string
	Name string
	MaterialCost string
//	BonusChance uint8
}
