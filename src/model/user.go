package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/rodericusifo/football-club-profile/libs/types"
	"gorm.io/gorm"
)

type User struct {
	ID        string           `json:"id" gorm:"primaryKey"`
	Username  string           `json:"username"`
	Email     string           `json:"email"`
	Password  string           `json:"password"`
	Address   types.NullString `json:"address"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	DeletedAt gorm.DeletedAt   `json:"deleted_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New().String()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}
