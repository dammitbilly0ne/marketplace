package recipe

type Repository interface {
	CreateRecipe(recipe *entities.Recipe) (*entities.Recipe, error)
}
