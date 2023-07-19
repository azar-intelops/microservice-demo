package daos

import (
	"errors"
	"github.com/azar-intelops/microservice-demo/alpha4/pkg/rest/server/daos/clients/sqls"
	"github.com/azar-intelops/microservice-demo/alpha4/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	InternalServerError = errors.New("internal server error")
)

type ColDao struct {
	db *gorm.DB
}

func NewColDao() (*ColDao, error) {
	sqlClient, err := sqls.InitGormMySQLDB()
	if err != nil {
		return nil, err
	}
	err = sqlClient.DB.AutoMigrate(models.Col{})
	if err != nil {
		return nil, err
	}
	return &ColDao{
		db: sqlClient.DB,
	}, nil
}

func (colDao *ColDao) CreateCol(m *models.Col) (*models.Col, error) {
	if err := colDao.db.Create(&m).Error; err != nil {
		return nil, err
	}

	log.Debugf("col created")
	return m, nil
}

func (colDao *ColDao) UpdateCol(id int64, m *models.Col) (*models.Col, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	var col *models.Col
	if err := colDao.db.Where("id = ?", id).First(&col).Error; err != nil {
		return nil, err
	}

	if id == col.Id {
		colDao.db.Save(&m)
		log.Debugf("col updated")
		return m, nil
	}

	return nil, InternalServerError
}

func (colDao *ColDao) DeleteCol(id int64) error {
	var m *models.Col
	if err := colDao.db.Where("id = ?", id).Delete(&m).Error; err != nil {
		return err
	}

	log.Debugf("col deleted")
	return nil
}

func (colDao *ColDao) ListCols() ([]*models.Col, error) {
	var cols []*models.Col
	if err := colDao.db.Find(&cols).Error; err != nil {
		return nil, err
	}

	log.Debugf("col listed")
	return cols, nil
}

func (colDao *ColDao) GetCol(id int64) (*models.Col, error) {
	var m *models.Col
	if err := colDao.db.Where("id = ?", id).First(&m).Error; err != nil {
		return nil, err
	}
	log.Debugf("col retrieved")
	return m, nil
}
