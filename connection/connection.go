package connection

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	User, Host, Psw, Name, Port string
	GormConfig                  gorm.Config
}

type Connection struct {
	*gorm.DB
}

func (c *Connection) Connect(config Config) (err error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.User, config.Host, config.Psw, config.Name, config.Port,
	)

	c.DB, err = gorm.Open(postgres.Open(dsn), &config.GormConfig)
	if err != nil {
		return
	}

	return
}
