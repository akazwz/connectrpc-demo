package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type RefreshToken struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	UserID    string    `json:"user_id"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (m *RefreshToken) TableName() string {
	return "refresh_tokens"
}

func (m *RefreshToken) BeforeCreate(tx *gorm.DB) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	m.ID = id.String()
	token := uuid.New().String()
	m.Token = token
	if m.ExpiresAt.IsZero() || m.ExpiresAt.Before(time.Now()) {
		m.ExpiresAt = time.Now().Add(time.Hour * 24 * 30)
	}
	return nil
}
