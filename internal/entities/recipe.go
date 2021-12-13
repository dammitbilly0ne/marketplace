package entities

type Recipe struct {
	ID     string
	Name   string
	Inputs []*MaterialCost
	Output []*RecipeOutput
}

type MaterialCost struct {
	ID   string
	Name string
	Qty  int
}

type RecipeOutput struct {
	ID   string
	Name string
	Qty  int
}

func (r *Recipe) GenerateOutput(actualInputs []*MaterialCost) []*RecipeOutput {
	for _, input := range r.Inputs {
		// verify the r.Inputs are here

	}
}
