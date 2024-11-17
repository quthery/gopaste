package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

type CodeNote struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Title      string    `json:"name" binding:"required"`
	Code       string    `json:"code" binding:"required"`
	Author     string    `json:"author" binding:"required"`
	AuthorID   uint      `json:"author_id" binding:"required"`
	Views      int       `gorm:"default:0" json:"views"`
	Stars      int       `gorm:"default:0" json:"stars"`
	Lang       string    `json:"lang" binding:"required"`
	Created_at time.Time `gorm:"autoCreateTime"`
	Updated_at time.Time `gorm:"autoUpdateTime"`
}

// TODO: Дополнить модели дополнительными полями
type User struct {
	ID         uint      `gorm:"primaryKey"`
	Username   string    `gorm:"unique" validate:"required"`
	Password   string    `json:"password" validate:"required"`
	Email      string    `gorm:"unique" validate:"required,email"`
	Created_at time.Time `gorm:"autoCreateTime"`
	Updated_at time.Time `gorm:"autoUpdateTime"`
	Deleted_at soft_delete.DeletedAt

	Containers []CodeNote `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;foreignKey:AuthorID;"`
}
