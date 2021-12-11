package entities

type Refined struct {
	ID           string
	Name         string
	Type         string
	Tier         string
	Rarity       string
	ResourceCost ResourceCost
}

type ResourceCost struct {
	ID       string
	Name     string
	Resource string
	Reagent  Reagent
}

type Reagent struct {
	ID     string
	Name   string
	Tier   string
	Rarity string
}
