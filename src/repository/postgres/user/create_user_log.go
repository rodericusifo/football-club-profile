package user_postgres_repository

import (
	"github.com/rodericusifo/football-club-profile/src/model"
)

func (r *UserRepository) CreateUserLog(payload *model.UserLog) error {
	model := new(model.UserLog)

	sql := r.db.Model(model)

	if payload != nil {
		model = payload
	}

	if err := sql.Create(model).Error; err != nil {
		return err
	}

	return nil
}
