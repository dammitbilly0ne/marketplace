package entities

type Refined struct {
	ID string
	Name string
	Type string
	Tier string
	Rarity string
	ResourceCost string
	}

type ResourceCost struct {
	 Resource string
	 Refined string
}