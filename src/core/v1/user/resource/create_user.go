package resource_user_v1

import (
	"github.com/rodericusifo/football-club-profile/src/model"
)

func (r *UserResource) CreateUser(payload *model.User) error {
	return r.Postgres.UserRepository.CreateUser(payload)
}
