package user

import (
	"gitlab.com/username/online-service-and-customer-care/entity"
)

//UserRepository specifies db opration for catagory
type UserRepository interface {
	Users() ([]entity.User, error)
	User(id uint) (entity.User, error)
	UpdateUser(user entity.User) error
	DeleteUser(id uint) error
	StoreUser(user entity.User) error
}
