package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"opendrive/utils"
)

type User struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	u.ID = id.String()
	if u.Password != "" {
		hashedPassword, err := utils.HashPassword(u.Password)
		if err != nil {
			return err
		}
		u.Password = hashedPassword
	}
	return nil
}

func (u *User) BeforeUpdate(tx *gorm.DB) error {
	if u.Password != "" {
		hashedPassword, err := utils.HashPassword(u.Password)
		if err != nil {
			return err
		}
		u.Password = hashedPassword
	}
	return nil
}
