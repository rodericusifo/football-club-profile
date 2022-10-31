package user_postgres_repository

import (
	"github.com/rodericusifo/football-club-profile/libs/types"
	"github.com/rodericusifo/football-club-profile/src/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(payload *model.User) error
	GetUser(query *types.Query, payload *model.User) (*model.User, error)
	CreateUserLog(payload *model.UserLog) error
}

type UserRepository struct {
	db *gorm.DB
}

func InitUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}
