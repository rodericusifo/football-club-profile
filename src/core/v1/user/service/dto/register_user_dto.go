package dto_service_user_v1

import (
	"github.com/rodericusifo/football-club-profile/libs/dto"
	request_controller_user_v1 "github.com/rodericusifo/football-club-profile/src/core/v1/user/controller/request"
)

type RegisterUserDTO struct {
	dto.DTO[any, any, *request_controller_user_v1.RegisterUserRequestBody]
}
