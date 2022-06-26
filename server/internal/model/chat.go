package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Chat struct {
	ID        string    `json:"id" gorm:"primaryKey"`
	Message   string    `json:"message" gorm:"not null;type:varchar(64)"`
	UserID    string    `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (p *Chat) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New().String()
	return nil
}
