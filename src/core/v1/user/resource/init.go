package resource_user_v1

import (
	"github.com/rodericusifo/football-club-profile/libs/database"
	"github.com/rodericusifo/football-club-profile/libs/types"
	"github.com/rodericusifo/football-club-profile/src/model"
	postgres_repository "github.com/rodericusifo/football-club-profile/src/repository/postgres"
)

type IUserResource interface {
	CreateUser(payload *model.User) error
	GetUser(query *types.Query, payload *model.User) (*model.User, error)
	CreateUserLog(payload *model.UserLog) error
}

type UserResource struct {
	Postgres *postgres_repository.PostgresRepository
}

func InitUserResource() IUserResource {
	var (
		postgres = database.InitPostgres()
	)

	return &UserResource{
		Postgres: postgres,
	}
}
