package recipe

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) CreateRecipe(recipe *entities.Recipe) (*entities.Recipe,error) {
	args := m.Called(recipe)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entities.Recipe), nil
}
