package entities

type Recipe struct {
	ID           string
	Name         string
	MaterialCost MaterialCost
}

type MaterialCost struct {
	ID       string
	Name     string
	Resource string
	Refined  string
	Reagent  Reagent
}
