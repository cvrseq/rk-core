package service

type StatementService interface {
	CreateConfig(cfg Model) (Model, error)
	GetUserById(id int) (*Model, error)
	GetUsersList() ([]Model, error)
	GetOrderById(orderId int) (*Model, error)
	GetOrdersList() ([]Model, error)
}

type VpnService struct {
	repo UserDataRepository
}

func NewVpnService(r UserDataRepository) StatementService {
	return &VpnService{repo: r}
}

func (s *VpnService) CreateConfig(cfg Model) error {
	if err := s.repo.CreateConfig(&cfg); err != nil {
		return Model{}, err
	}
	return cfg, nil
}
