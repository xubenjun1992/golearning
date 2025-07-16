package service

import (
	"blog/dao"
	"blog/utils"
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(username, password string) (string, error)
}
type authService struct {
	userDao dao.UserDao
}

func NewAuthService(u dao.UserDao) AuthService {
	return &authService{
		userDao: u,
	}
}

func (a *authService) Login(username, password string) (string, error) {
	user, error := a.userDao.GetUserByName(username)

	if error != nil {
		return "", error
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		log.Printf("Password comparison failed: %v", err)
		return "", errors.New("invalid username or password")
	}
	token := utils.GenerateToken(int64(user.Id), user.Username)
	return token, nil
}
