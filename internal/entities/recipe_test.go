package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupFixture() *Recipe {
	return &Recipe{
		ID:   "1234",
		Name: "timber",
		Inputs: MaterialCost{
			ID:   "7890",
			Name: "green_wood",
			Qty:  2,
		},
	}
}

func TestRecipe_GenerateOutput(t *testing.T) {
	t.Run("it requires the correct input", func(t *testing.T) {

	})
	t.Run("it calculates correctly", func(t *testing.T) {
		r := setupFixture()

		results := r.GenerateOutput()

		assert.Equal(t, "timber", results[0].Name)
	})
}
