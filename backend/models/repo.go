package models

import (
	"time"

	"gorm.io/gorm"
)

type Repo struct {
	gorm.Model
	Name          string `gorm:"not null"`
	Description   string
	URL           string `gorm:"not null;uniqueIndex"`
	LastCheckedAt time.Time

	Users  []User  `gorm:"many2many:user_favorite_repos;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Issues []Issue `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
