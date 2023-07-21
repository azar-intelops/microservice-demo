package daos

import (
	"errors"
	"github.com/azar-intelops/microservice-demo/beta-gormsqlite/pkg/rest/server/daos/clients/sqls"
	"github.com/azar-intelops/microservice-demo/beta-gormsqlite/pkg/rest/server/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao() (*UserDao, error) {
	sqlClient, err := sqls.InitGormSqliteDB()
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
		log.Errorf("failed to create user: %v", err)
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
		log.Errorf("failed to find user for update: %v", err)
		return nil, err
	}

	if err := userDao.db.Save(&m).Error; err != nil {
		log.Errorf("failed to update user: %v", err)
		return nil, err
	}
	log.Debugf("user updated")
	return m, nil
}

func (userDao *UserDao) DeleteUser(id int64) error {
	var m *models.User
	if err := userDao.db.Where("id = ?", id).Delete(&m).Error; err != nil {
		log.Errorf("failed to delete user: %v", err)
		return err
	}

	log.Debugf("user deleted")
	return nil
}

func (userDao *UserDao) ListUsers() ([]*models.User, error) {
	var users []*models.User
	if err := userDao.db.Find(&users).Error; err != nil {
		log.Errorf("failed to list users: %v", err)
		return nil, err
	}

	log.Debugf("user listed")
	return users, nil
}

func (userDao *UserDao) GetUser(id int64) (*models.User, error) {
	var m *models.User
	if err := userDao.db.Where("id = ?", id).First(&m).Error; err != nil {
		log.Errorf("failed to get user: %v", err)
		return nil, err
	}
	log.Debugf("user retrieved")
	return m, nil
}
