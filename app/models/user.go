package models

import "time"

type UserModel struct {
	ID           uint    `gorm:"primary_key"`
	Name     string  `gorm:"column:name"`
	Email        string  `gorm:"column:email;unique_index"`
	Bio          string  `gorm:"column:bio;size:1024"`
	Image        *string `gorm:"column:image"`
	Password string  `gorm:"column:password;not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}
