package models

type RequestsLocation struct {
	RequestID  string `gorm:"type:uuid;primary_key"`
	LocationID string `gorm:"type:varchar(10);primary_key"`
}

func (r RequestsLocation) TableName() string {
	return "pesquisai.request_locations"
}
