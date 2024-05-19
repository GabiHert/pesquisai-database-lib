package models

type Location struct {
	ID string `gorm:"type:varchar(10);primary_key;"`
}

func (r Location) TableName() string {
	return "pesquisai.locations"
}
