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

func (c *UserController) LoginUser(ctx echo.Context) error {
	reqBody := new(request_controller_user_v1.LoginUserRequestBody)
	if err := validator.ValidateRequest(ctx, reqBody); err != nil {
		return err
	}

	userWithTokenDto, err := c.UserService.LoginUser(&dto_service_user_v1.LoginUserDTO{
		DTO: dto.DTO[any, any, *request_controller_user_v1.LoginUserRequestBody]{
			Param: nil,
			Query: nil,
			Body:  reqBody,
		},
	})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response.ResponseSuccess("login user success", userWithTokenDto))
}
