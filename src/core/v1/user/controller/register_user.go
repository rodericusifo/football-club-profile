package controller_user_v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rodericusifo/football-club-profile/libs/dto"
	"github.com/rodericusifo/football-club-profile/libs/response"
	"github.com/rodericusifo/football-club-profile/libs/validator"
	request_controller_user_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/user/controller/request"
	dto_service_user_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/user/service/dto"
)

func (c *UserController) RegisterUser(ctx echo.Context) error {
	reqBody := new(request_controller_user_v1.RegisterUserRequestBody)
	if err := validator.ValidateRequest(ctx, reqBody); err != nil {
		return err
	}

	if err := c.UserService.RegisterUser(&dto_service_user_v1.RegisterUserDTO{
		DTO: dto.DTO[any, any, *request_controller_user_v1.RegisterUserRequestBody]{
			Param: nil,
			Query: nil,
			Body:  reqBody,
		},
	}); err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, response.ResponseSuccess[any]("success register user", nil))
}
