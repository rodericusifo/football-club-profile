package service_user_v1

import (
	resource_user_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/user/resource"
	dto_service_user_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/user/service/dto"
)

type IUserService interface {
	RegisterUser(payload *dto_service_user_v1.RegisterUserDTO) error
	LoginUser(payload *dto_service_user_v1.LoginUserDTO) (*dto_service_user_v1.UserWithTokenDTO, error)
}

type UserService struct {
	UserResource resource_user_v1.IUserResource
}

func InitUserService(userResource resource_user_v1.IUserResource) IUserService {
	return &UserService{
		UserResource: userResource,
	}
}
