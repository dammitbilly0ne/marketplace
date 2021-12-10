package entities

type MaterialCost struct {
	Resource string
	Refined string
	Reagent string
}

type Recipe struct {
	ID string
	Name string
	MaterialCost string
}
