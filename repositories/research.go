package repositories

import (
	"github.com/GabiHert/pesquisai-database-lib/connection"
	"github.com/GabiHert/pesquisai-database-lib/models"
)

type Research struct {
	Connection *connection.Connection
}

func (s *Research) Create(research *models.Research) error {
	res := s.Connection.Create(research)
	return res.Error
}

func (s *Research) UpdateStatus(id, status string) error {
	res := s.Connection.Model(&models.Research{ID: id}).UpdateColumn("status", status)
	return res.Error
}

func (s *Research) Get(id string) (research models.Research, err error) {
	res := s.Connection.First(&research, id)
	return research, res.Error
}
