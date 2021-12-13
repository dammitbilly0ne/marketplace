package item

import "github.com/dammitbilly0ne/marketplace/internal/entities"

type Repository interface {
	CreateItem(item *entities.Item) (*entities.Item, error)
}
