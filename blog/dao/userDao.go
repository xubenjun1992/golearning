package dao

import (
	"blog/core"
	"blog/model"
)

type UserDao interface {
	CreateUser(user model.User) (*model.User, error)
	GetUserById(id uint) (*model.User, error)
	GetUserByName(name string) (*model.User, error)
}

type userDao struct{}

func NewUserDao() UserDao {
	return &userDao{}
}

func (u *userDao) GetUserById(id uint) (*model.User, error) {
	db := core.GetDb()
	var user model.User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (u *userDao) CreateUser(user model.User) (*model.User, error) {
	db := core.GetDb()
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (u *userDao) GetUserByName(name string) (*model.User, error) {
	db := core.GetDb()
	var user model.User
	if err := db.Where("username = ?", name).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
