package repositories

import (
	"context"
	"github.com/GabiHert/pesquisai-database-lib/connection"
	"github.com/GabiHert/pesquisai-database-lib/models"
)

type Research struct {
	Connection *connection.Connection
}

func (s *Research) Create(ctx context.Context, research *models.Research) error {
	res := s.Connection.WithContext(ctx).Create(research)
	return res.Error
}

func (s *Research) UpdateStatus(ctx context.Context, id, status string) error {
	res := s.Connection.WithContext(ctx).Model(&models.Research{ID: &id}).UpdateColumn("status", status)
	return res.Error
}

func (s *Research) Get(ctx context.Context, id string) (research *models.Research, err error) {
	res := s.Connection.WithContext(ctx).First(&research, id)
	return research, res.Error
}
