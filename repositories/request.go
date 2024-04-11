package repositories

import (
	"github.com/GabiHert/pesquisai-database-lib/connection"
	"github.com/GabiHert/pesquisai-database-lib/models"
)

type Request struct {
	connection *connection.Connection
}

func (s *Request) Create(request *models.Request) error {
	res := s.connection.Create(request)
	return res.Error
}

func (s *Request) GetTotals(id string) (total, totalFinishedResearched int, err error) {
	var request models.Request
	res := s.connection.Select("total_researches", "total_finished_researches").First(&request, id)
	return request.TotalResearches, request.TotalFinishedResearched, res.Error
}

func (s *Request) Get(id string) (request models.Request, err error) {
	res := s.connection.First(&request, id)
	return request, res.Error
}

func (s *Request) GetWithResearches(id string) (request models.Request, err error) {
	res := s.connection.Preload("Researches").First(&request, id)
	return request, res.Error
}
