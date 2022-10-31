package controller_user_v1

import (
	"github.com/labstack/echo/v4"
	service_user_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/user/service"
)

type UserController struct {
	UserService service_user_v1.IUserService
}

func InitUserController(userService service_user_v1.IUserService) *UserController {
	return &UserController{UserService: userService}
}

func (userController *UserController) Mount(group *echo.Group) {
	group.POST("/register", userController.RegisterUser)
	group.POST("/login", userController.LoginUser)
}
