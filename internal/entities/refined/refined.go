package refined

type Refined struct {
	ID string
	Name string
	Type string
	Tier string
	ResourceCost string
//	BonusChance string    or    *int (decimal numbers / calculations??
	}

type ResourceCost struct {
	 Resources string
	 Refined string
	 Reagent string
}