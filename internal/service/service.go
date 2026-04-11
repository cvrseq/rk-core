package service

type StatementService interface {
	CreateConfig(cfg Model) (*Model, error)
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

func (s *VpnService) CreateConfig(cfg Model) (*Model, error) {

	return nil, nil
}

func (s *VpnService) GetUserById(id int) (*Model, error) {
	return nil, nil
}

func (s *VpnService) GetUsersList() ([]Model, error) {
	return nil, nil
}

func (s *VpnService) GetOrderById(orderId int) (*Model, error) {
	return nil, nil
}

func (s *VpnService) GetOrdersList() ([]Model, error) {
	return nil, nil
}
