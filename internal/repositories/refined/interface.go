package refined

import "github.com/dammitbilly0ne/marketplace/internal/entities"

type Repository interface {
	CreateRefined(refined *entities.Refined) (*entities.Refined, error)
}
