package services

import (
	"github.com/azar-intelops/microservice-demo/test-service/pkg/rest/server/daos"
	"github.com/azar-intelops/microservice-demo/test-service/pkg/rest/server/models"
)

type UserToLoginService struct {
	userToLoginDao *daos.UserToLoginDao
}

func NewUserToLoginService() (*UserToLoginService, error) {
	userToLoginDao, err := daos.NewUserToLoginDao()
	if err != nil {
		return nil, err
	}
	return &UserToLoginService{
		userToLoginDao: userToLoginDao,
	}, nil
}

func (userToLoginService *UserToLoginService) CreateUserToLogin(userToLogin *models.UserToLogin) (*models.UserToLogin, error) {
	return userToLoginService.userToLoginDao.CreateUserToLogin(userToLogin)
}

func (userToLoginService *UserToLoginService) UpdateUserToLogin(id int64, userToLogin *models.UserToLogin) (*models.UserToLogin, error) {
	return userToLoginService.userToLoginDao.UpdateUserToLogin(id, userToLogin)
}

func (userToLoginService *UserToLoginService) DeleteUserToLogin(id int64) error {
	return userToLoginService.userToLoginDao.DeleteUserToLogin(id)
}

func (userToLoginService *UserToLoginService) ListUserToLogins() ([]*models.UserToLogin, error) {
	return userToLoginService.userToLoginDao.ListUserToLogins()
}

func (userToLoginService *UserToLoginService) GetUserToLogin(id int64) (*models.UserToLogin, error) {
	return userToLoginService.userToLoginDao.GetUserToLogin(id)
}
