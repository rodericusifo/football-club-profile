package service_user_v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rodericusifo/football-club-profile/libs/util"
	request_controller_user_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/user/controller/request"
	dto_service_user_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/user/service/dto"
	"github.com/rodericusifo/football-club-profile/src/model"
)

func (s *UserService) RegisterUser(payload *dto_service_user_v1.RegisterUserDTO) error {
	userModelRes, err := s.UserResource.GetUser(nil, &model.User{
		Username: payload.Body.Username,
	})
	if err != nil {
		if err.Error() != "record not found" {
			return err
		}
	}
	if userModelRes != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "user already registered")
	}

	userModel, err := util.ConvertToStruct[model.User](struct {
		*request_controller_user_v1.RegisterUserRequestBody
	}{payload.Body})
	if err != nil {
		return err
	}

	hashedPassword, err := util.HashPassword(userModel.Password)
	if err != nil {
		return err
	}
	userModel.Password = hashedPassword

	err = s.UserResource.CreateUser(&userModel)
	if err != nil {
		return err
	}

	return nil
}
