package models

import (
	"time"
)

type Research struct {
	ID        *string   `gorm:"type:uuid;primaryKey"`
	RequestID *string   `gorm:"type:uuid"`
	Title     *string   `gorm:"type:varchar(50)"`
	Link      *string   `gorm:"type:varchar(100)"`
	Status    *string   `gorm:"type:varchar(10)"`
	Summary   *string   `gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time
}
