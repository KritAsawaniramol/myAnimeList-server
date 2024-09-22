package entities

import (
	"time"

	"gorm.io/gorm"
)

type (
	User struct {
		gorm.Model
		Name               string
		Email              string
		AvatarURL          string
		AuthProviderName   string
		AuthProviderUserID string
		Credential         []Credential
		AnimeLists         []AnimeLists
		Comments           []Comments
	}

	Credential struct {
		gorm.Model
		UserID       uint   `gorm:"not null"`
		AccessToken  string `gorm:"not null"`
		RefreshToken string `gorm:"not null"`
	}

	AnimeLists struct {
		UserID        uint   `gorm:"primaryKey;not null"`
		MalID         string `gorm:"primaryKey;not null;check:mal_id <> ''"`
		Status        string `gorm:"not null"`
		EpisodesCount uint   `gorm:"not null"`
		Score         int    `gorm:"not null"`
		CreatedAt     time.Time
		UpdatedAt     time.Time
	}

	Comments struct {
		gorm.Model
		UserID   uint `gorm:"not null"`
		User     User
		MalID    string `gorm:"not null;check:mal_id <> ''"`
		ParentID uint
		Content  string
	}
)
