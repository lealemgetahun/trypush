package service

import (
	"gitlab.com/username/online-service-and-customer-care/entity"
	"gitlab.com/username/online-service-and-customer-care/user"
)

// CommentService implements menu.CommentService interface
type UserService struct {
	userRepo user.UserRepository
}

// NewCommentService returns a new CommentService object
func NewUserService(userRepo user.UserRepository) user.UserService {
	return &UserService{userRepo: userRepo}
}

// Users returns all stored users
func (us *UserService) Users() ([]entity.User, []error) {
	usrs, errs := us.userRepo.Users()
	if len(errs) > 0 {
		return nil, errs
	}
	return usrs, errs
}

// User retrieves stored User by its id
func (us *UserService) User(id uint) (*entity.User, []error) {
	usr, errs := us.userRepo.User(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// UpdateUser updates a given user
func (us *UserService) UpdateUser(user *entity.User) (*entity.User, []error) {
	usr, errs := us.userRepo.UpdateUser(user)
	if len(errs) > 0 {
		return nil, errs
	}
	return usr, errs
}

// DeleteUser deletes a given user
func (us *UserService) DeleteUser(id uint) (*entity.User, []error) {
	user, errs := us.userRepo.DeleteUser(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return user, errs
}

// StoreUser stores a given user
func (us *UserService) StoreUser(user *entity.User) (*entity.User, []error) {
	return nil, nil
}
