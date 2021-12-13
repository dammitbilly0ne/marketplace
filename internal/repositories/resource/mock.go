package resource

import (
	"github.com/dammitbilly0ne/marketplace/internal/entities"
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) CreateResource(resource *entities.Resource) (*entities.Resource, error) {
	args := m.Called(resource)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entities.Resource), nil
}
