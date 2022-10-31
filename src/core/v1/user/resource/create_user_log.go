package resource_user_v1

import (
	"github.com/rodericusifo/football-club-profile/src/model"
)

func (r *UserResource) CreateUserLog(payload *model.UserLog) error {
	return r.Postgres.UserRepository.CreateUserLog(payload)
}
