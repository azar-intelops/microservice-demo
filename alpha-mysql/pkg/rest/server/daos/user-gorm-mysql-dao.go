package daos

import (
	"errors"
	"github.com/azar-intelops/microservice-demo/alpha-mysql/pkg/rest/server/daos/clients/sqls"
	"github.com/azar-intelops/microservice-demo/alpha-mysql/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	InternalServerError = errors.New("internal server error")
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao() (*UserDao, error) {
	sqlClient, err := sqls.InitGormMySQLDB()
	if err != nil {
		return nil, err
	}
	err = sqlClient.DB.AutoMigrate(models.User{})
	if err != nil {
		return nil, err
	}
	return &UserDao{
		db: sqlClient.DB,
	}, nil
}

func (userDao *UserDao) CreateUser(m *models.User) (*models.User, error) {
	if err := userDao.db.Create(&m).Error; err != nil {
		return nil, err
	}

	log.Debugf("user created")
	return m, nil
}

func (userDao *UserDao) UpdateUser(id int64, m *models.User) (*models.User, error) {
	if id == 0 {
		return nil, errors.New("invalid updated ID")
	}
	if id != m.Id {
		return nil, errors.New("id and payload don't match")
	}

	var user *models.User
	if err := userDao.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	if id == user.Id {
		userDao.db.Save(&m)
		log.Debugf("user updated")
		return m, nil
	}

	return nil, InternalServerError
}

func (userDao *UserDao) DeleteUser(id int64) error {
	var m *models.User
	if err := userDao.db.Where("id = ?", id).Delete(&m).Error; err != nil {
		return err
	}

	log.Debugf("user deleted")
	return nil
}

func (userDao *UserDao) ListUsers() ([]*models.User, error) {
	var users []*models.User
	if err := userDao.db.Find(&users).Error; err != nil {
		return nil, err
	}

	log.Debugf("user listed")
	return users, nil
}

func (userDao *UserDao) GetUser(id int64) (*models.User, error) {
	var m *models.User
	if err := userDao.db.Where("id = ?", id).First(&m).Error; err != nil {
		return nil, err
	}
	log.Debugf("user retrieved")
	return m, nil
}
