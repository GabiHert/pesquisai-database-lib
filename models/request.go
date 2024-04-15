package models

import (
	"time"
)

type Request struct {
	ID                      *string    `gorm:"type:uuid;primaryKey"`
	Context                 *string    `gorm:"type:varchar(1000)"`
	Research                *string    `gorm:"type:varchar(1000)"`
	TotalResearches         *int       `gorm:"default:0"`
	TotalFinishedResearches *int       `gorm:"default:0"`
	Status                  *string    `gorm:"type:varchar(10);default:'PENDING'"`
	Overall                 *string    `gorm:"type:text"`
	CreatedAt               *time.Time `gorm:"autoCreateTime"`
	UpdatedAt               *time.Time
	Researches              []Research `gorm:"foreignKey:RequestID"`
}
