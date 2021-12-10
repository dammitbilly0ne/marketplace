package item

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) CreateItem(item *entities.Item) (*entities.Item,error) {
	args := m.Called(item)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entities.Item), nil
}
