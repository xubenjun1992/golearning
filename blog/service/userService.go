package service

import (
	"blog/dao"
	"blog/model"
	"log"
)

type UserService interface {
	CreateUser(user model.User) (*model.User, error)
	GetUserById(id uint) (*model.User, error)
	GetUserByName(name string) (*model.User, error)
}

type userService struct {
	userDao dao.UserDao
}

func NewUserService(u dao.UserDao) UserService {
	return &userService{
		userDao: u,
	}
}
func (u *userService) CreateUser(user model.User) (*model.User, error) {
	if user, err := u.userDao.CreateUser(user); err != nil {
		return nil, err
	} else {
		return user, nil
	}
}
func (u *userService) GetUserById(id uint) (*model.User, error) {
	if user, err := u.userDao.GetUserById(id); err != nil {
		return nil, err
	} else {
		return user, nil
	}
}
func (u *userService) GetUserByName(name string) (*model.User, error) {
	if user, err := u.userDao.GetUserByName(name); err != nil {
		log.Printf("Error getting user by name: %v", err)
		return nil, err
	} else {
		return user, nil
	}
}
