package resource_user_v1

import (
	"github.com/rodericusifo/football-club-profile/libs/types"
	"github.com/rodericusifo/football-club-profile/src/model"
)

func (r *UserResource) GetUser(query *types.Query, payload *model.User) (*model.User, error) {
	return r.Postgres.UserRepository.GetUser(query, payload)
}
