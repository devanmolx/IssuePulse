package models

import "gorm.io/gorm"

type Issue struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	Status      string `gorm:"not null;default:'open'"`
	RepoID      uint   `gorm:"not null;index"`
	Repo        Repo   `gorm:"foreignKey:RepoID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
