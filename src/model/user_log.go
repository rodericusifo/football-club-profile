package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserLog struct {
	ID          string         `json:"id" gorm:"primaryKey"`
	Username    string         `json:"name"`
	Password    string         `json:"password"`
	IsValid     bool           `json:"is_valid"`
	AccessToken string         `json:"access_token"`
	LoginAt     time.Time      `json:"login_at"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

func (u *UserLog) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New().String()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}
