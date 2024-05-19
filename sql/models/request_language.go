package models

type RequestsLanguage struct {
	RequestID  string `gorm:"type:uuid;primary_key"`
	LanguageID string `gorm:"type:varchar(10);primary_key"`
}

func (r RequestsLanguage) TableName() string {
	return "pesquisai.request_languages"
}
