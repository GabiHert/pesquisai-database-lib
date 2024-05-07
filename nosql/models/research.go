package models

import (
	"time"
)

type Research struct {
	ID        *string    `bson:"_id,omitempty"`
	RequestID *string    `bson:"requestID,omitempty"`
	Title     *string    `bson:"title,omitempty"`
	Link      *string    `bson:"link,omitempty"`
	Status    *string    `bson:"status,omitempty"`
	Summary   *string    `bson:"summary,omitempty"`
	CreatedAt *time.Time `bson:"createdAt,omitempty"`
	UpdatedAt *time.Time `bson:"updatedAt,omitempty"`
}
