package models

type Language struct {
	ID string `gorm:"type:varchar(10);primary_key;"`
}

func (r Language) TableName() string {
	return "pesquisai.languages"
}
