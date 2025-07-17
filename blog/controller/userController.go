package controller

import (
	"blog/model"
	"blog/reqdto"
	"blog/service"
	"blog/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserController interface {
	RegistUser(c *gin.Context)
	GetUserById(c *gin.Context)
	Login(c *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &userController{
		userService: userService,
	}
}

func (u *userController) RegistUser(ctx *gin.Context) {
	var createUserDTO reqdto.CreateUserDTO
	if err := ctx.ShouldBindJSON(&createUserDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(createUserDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user := model.User{
		Username: createUserDTO.Username,
		Password: string(hashPassword),
		Email:    createUserDTO.Email,
	}
	if createdUser, err := u.userService.CreateUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": createdUser})
	}

}

func (u *userController) GetUserById(ctx *gin.Context) {
	userId := ctx.GetUint("userId")
	if user, err := u.userService.GetUserById(userId); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"user": user})
	}
}

func (u *userController) Login(ctx *gin.Context) {
	var loginDTO reqdto.CreateUserDTO
	if err := ctx.ShouldBindJSON(&loginDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := u.userService.GetUserByName(loginDTO.Username)
	if err != nil || user == nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDTO.Password)); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	ctx.Set("userId", user.Id)
	token := utils.GenerateToken(int64(user.Id), user.Username)
	ctx.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
