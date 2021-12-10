package resource

type Repository interface {
	CreateResource(resource *entities.Resource) (*entities.Resource, error)
}