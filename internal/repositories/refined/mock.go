package refined

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) CreateRefined(refined *entities.Refined) (*entities.Refined,error) {
	args := m.Called(refined)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*entities.Refined), nil
}