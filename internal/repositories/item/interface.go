package item

type Repository interface {
	CreateItem(item *entities.Item) (*entities.Item, error)
}
