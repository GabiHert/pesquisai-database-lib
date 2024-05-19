package repositories

import (
	"context"
	"github.com/PesquisAi/pesquisai-database-lib/sql/connection"
	"github.com/PesquisAi/pesquisai-database-lib/sql/models"
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

func (s *Request) GetWithRelations(ctx context.Context, id string) (request *models.Request, err error) {
	res := s.Connection.WithContext(ctx).
		Preload("Locations").Preload("Languages").Preload("Researches").
		First(&request, "id = ?", id)
	return request, res.Error
}

func (s *Request) UpdateStatus(ctx context.Context, id, status string) error {
	res := s.Connection.WithContext(ctx).Model(&models.Request{ID: &id}).UpdateColumn("status", status)
	return res.Error
}

func (s *Request) RelateLanguage(ctx context.Context, id string, language string) error {
	return s.Connection.WithContext(ctx).Create(&models.RequestsLanguage{
		RequestID:  id,
		LanguageID: language,
	}).Error
}

func (s *Request) RelateLocation(ctx context.Context, id string, location string) error {
	return s.Connection.WithContext(ctx).Create(&models.RequestsLocation{
		RequestID:  id,
		LocationID: location,
	}).Error
}
