package user

import (
	"gitlab.com/username/online-service-and-customer-care/entity"
)

// UserService imp
type UserService interface {
	Users() ([]entity.User, error)
	User(id uint) (*entity.User, error)
	UpdateUser(user *entity.User) error
	DeleteUser(id uint) (*entity.User, error)
	StoreUser(user *entity.User) error
}
