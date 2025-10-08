package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username        string `gorm:"not null"`
	Email           string `gorm:"uniqueIndex;not null"`
	GithubID        string `gorm:"uniqueIndex;not null"`
	ProfileImageURL string
	AccessToken     string `gorm:"not null"`
	LastSeenAt      *time.Time

	FavoriteRepos []Repo `gorm:"many2many:user_favorite_repos;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
