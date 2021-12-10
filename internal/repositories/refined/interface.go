package refined

type Repository interface {
	CreateRefined(refined *entities.Refined) (*entities.Refined, error)
}