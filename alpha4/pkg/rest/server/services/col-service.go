package services

import (
	"github.com/azar-intelops/microservice-demo/alpha4/pkg/rest/server/daos"
	"github.com/azar-intelops/microservice-demo/alpha4/pkg/rest/server/models"
)

type ColService struct {
	colDao *daos.ColDao
}

func NewColService() (*ColService, error) {
	colDao, err := daos.NewColDao()
	if err != nil {
		return nil, err
	}
	return &ColService{
		colDao: colDao,
	}, nil
}

func (colService *ColService) CreateCol(col *models.Col) (*models.Col, error) {
	return colService.colDao.CreateCol(col)
}

func (colService *ColService) UpdateCol(id int64, col *models.Col) (*models.Col, error) {
	return colService.colDao.UpdateCol(id, col)
}

func (colService *ColService) DeleteCol(id int64) error {
	return colService.colDao.DeleteCol(id)
}

func (colService *ColService) ListCols() ([]*models.Col, error) {
	return colService.colDao.ListCols()
}

func (colService *ColService) GetCol(id int64) (*models.Col, error) {
	return colService.colDao.GetCol(id)
}
