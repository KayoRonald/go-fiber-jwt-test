package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func (u *User) BeforeSave() {
	u.ID = uuid.New()
}
