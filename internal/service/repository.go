package service

type UserDataRepository interface {
	GetUserById(id int) (*Model, error)
	GetUsersList() ([]Model, error)
	GetOrderById(orderId int) (*Model, error)
	GetOrdersList() ([]Model, error) 
}
