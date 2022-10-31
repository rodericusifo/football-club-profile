package service_user_v1

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/rodericusifo/football-club-profile/libs/config"
	"github.com/rodericusifo/football-club-profile/libs/types"
	"github.com/rodericusifo/football-club-profile/libs/util"
	dto_service_user_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/user/service/dto"
	"github.com/rodericusifo/football-club-profile/src/model"
)

func (s *UserService) LoginUser(payload *dto_service_user_v1.LoginUserDTO) (*dto_service_user_v1.UserWithTokenDTO, error) {
	userModelRes, err := s.UserResource.GetUser(nil, &model.User{
		Username: payload.Body.Username,
	})
	if err != nil {
		return nil, err
	}

	match := util.CheckPasswordHash(payload.Body.Password, userModelRes.Password)
	if !match {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "username and password not match")
	}

	claims := &types.JwtCustomClaims{
		ID:       userModelRes.ID,
		Username: userModelRes.Username,
		Email:    userModelRes.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.AppConfig.JWT.SecretKey))
	if err != nil {
		return nil, err
	}

	userWithTokenDto, err := util.ConvertToStruct[*dto_service_user_v1.UserWithTokenDTO](&dto_service_user_v1.UserWithTokenDTO{
		User: dto_service_user_v1.UserDTOForLogin{
			ID: userModelRes.ID,
		},
		Token: t,
	})
	if err != nil {
		return nil, err
	}

	err = s.UserResource.CreateUserLog(&model.UserLog{
		Username:    userModelRes.Username,
		Password:    userModelRes.Password,
		IsValid:     true,
		AccessToken: t,
		LoginAt:     time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return userWithTokenDto, nil
}
