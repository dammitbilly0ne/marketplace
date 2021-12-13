package recipe

import "github.com/dammitbilly0ne/marketplace/internal/entities"

type Repository interface {
	CreateRecipe(recipe *entities.Recipe) (*entities.Recipe, error)
}
