package models

import (
	"time"
)

type Operation struct {
	ID            uint              `json:"id" gorm:"primaryKey"`
	OperationType OperationType     `json:"-" gorm:"foreignKey:TypeID;references:ID"`
	TypeID        uint              `json:"type_id"`
	Category      OperationCategory `json:"-" gorm:"foreignKey:CategoryID;references:ID"`
	CategoryID    uint              `json:"category_id"`
	Amount        float64           `json:"amount"`
	Description   string            `json:"description"`
	User          User              `json:"-" gorm:"foreignKey:UserID;references:ID"`
	UserID        uint              `json:"-"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"-"`
	IsDeleted     bool              `json:"-" gorm:"default:false"`
}

type OperationType struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"unique" json:"title"`
	IsActive  bool      `gorm:"default:false" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type OperationCategory struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `gorm:"unique" json:"title"`
	IsActive  bool      `gorm:"default:false" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
