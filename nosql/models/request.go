package models

import (
	"time"
)

type Request struct {
	ID        *string    `bson:"_id,omitempty"`
	Context   *string    `bson:"context,omitempty"`
	Research  *string    `bson:"research,omitempty"`
	Status    *string    `bson:"status,omitempty"`
	Languages *[]string  `bson:"languages,omitempty"`
	Locations *[]string  `bson:"locations,omitempty"`
	Sentences *[]string  `bson:"sentences,omitempty"`
	CreatedAt *time.Time `bson:"createdAt,omitempty"`
	UpdatedAt *time.Time `bson:"updatedAt,omitempty"`
}
