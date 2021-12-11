package resource

import "github.com/dammitbilly0ne/marketplace/internal/entities"

type Repository interface {
	CreateResource(resource *entities.Resource) (*entities.Resource, error)
}