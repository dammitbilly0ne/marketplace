package handlers

import (
	"errors"
	"testing"
)

func setupFixture() *Alpha {
	return &Alpha{
		itemRepo: &items.MockRepo{},
		recipeRepo: &recipe.MockRepo{},
		refinedRepo: &refined.MockRepo{},
		resourceRepo: &resource.MockRepo{},
	}
}

func TestNewAlpha(t *testing.T) {
	t.Run("it requires a config", func(t *testing.T) {
		actual, err := NewAlpha(nil)
		expErr := errors.New(configReqMsg)

		assert.Nil(t, actual)
		assert.NotNil (t, err)
		assert.Equal (t, expErr, err)
	})

	t.Run("it requires an itemRepo", func(t *testing.T) {
		actual, err := NewAlpha(&AlphaConfig{
			ItemRepo: nil,
		})
		expErr := errors.New(itemRepoReqMsg)

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expErr, err)
	})

	t.Run("it requires an recipeRepo", func(t *testing.T) {
		actual, err := NewAlpha(&AlphaConfig{
			ItemRepo: nil,
		})
		expErr := errors.New(recipeRepoReqMsg)

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expErr, err)
	})

	t.Run("it requires an refinedRepo", func(t *testing.T) {
		actual, err := NewAlpha(&AlphaConfig{
			ItemRepo: nil,
		})
		expErr := errors.New(refinedRepoReqMsg)

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expErr, err)
	})

	t.Run("it requires an resourceRepo", func(t *testing.T) {
		actual, err := NewAlpha(&AlphaConfig{
			ItemRepo: nil,
		})
		expErr := errors.New(resourceRepoReqMsg)

		assert.Nil(t, actual)
		assert.NotNil(t, err)
		assert.Equal(t, expErr, err)
	})

	t.Run("it returns an alpha", func(t *testing.T) {
		actual, err := NewAlpha(&AlphaConfig{
			ItemRepo: &items.MockRepo{},
			RecipeRepo: &recipes.MockRepo{},
			RefinedRepo: &refined.MockRepo{},
			ResourceRepo: &resoures.MockRepo{},
		})

		assert.NotNil(t, actual)
		assert.Nil(t,err)
		assertEqual(t, actual.ItemRepo)
	})

}