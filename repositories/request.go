package repositories

import (
	"context"
	"github.com/PesquisAi/pesquisai-database-lib/connection"
	"github.com/PesquisAi/pesquisai-database-lib/models"
)

type Request struct {
	Connection *connection.Connection
}

func (s *Request) Create(ctx context.Context, request *models.Request) error {
	res := s.Connection.WithContext(ctx).Create(request)
	return res.Error
}

func (s *Request) GetTotals(ctx context.Context, id string) (total, totalFinishedResearched *int, err error) {
	var request models.Request

	res := s.Connection.WithContext(ctx).Select("total_researches", "total_finished_researches").First(&request, id)
	return request.TotalResearches, request.TotalFinishedResearches, res.Error
}

func (s *Request) Get(ctx context.Context, id string) (request *models.Request, err error) {
	res := s.Connection.WithContext(ctx).First(&request, id)
	return request, res.Error
}
func (s *Request) GetAttributes(ctx context.Context, id string, attributes ...string) (request *models.Request, err error) {
	res := s.Connection.WithContext(ctx).Select(attributes).First(&request, id)
	return request, res.Error
}

func (s *Request) GetWithResearches(ctx context.Context, id string) (request *models.Request, err error) {
	res := s.Connection.WithContext(ctx).Preload("Researches").First(&request, id)
	return request, res.Error
}
