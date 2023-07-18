package controllers

import (
	"github.com/azar-intelops/microservice-demo/test-service/pkg/rest/server/models"
	"github.com/azar-intelops/microservice-demo/test-service/pkg/rest/server/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type UserToLoginController struct {
	userToLoginService *services.UserToLoginService
}

func NewUserToLoginController() (*UserToLoginController, error) {
	userToLoginService, err := services.NewUserToLoginService()
	if err != nil {
		return nil, err
	}
	return &UserToLoginController{
		userToLoginService: userToLoginService,
	}, nil
}

func (userToLoginController *UserToLoginController) CreateUserToLogin(context *gin.Context) {
	// validate input
	var input models.UserToLogin
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger userToLogin creation
	if _, err := userToLoginController.userToLoginService.CreateUserToLogin(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "UserToLogin created successfully"})
}

func (userToLoginController *UserToLoginController) UpdateUserToLogin(context *gin.Context) {
	// validate input
	var input models.UserToLogin
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger userToLogin update
	if _, err := userToLoginController.userToLoginService.UpdateUserToLogin(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "UserToLogin updated successfully"})
}

func (userToLoginController *UserToLoginController) FetchUserToLogin(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger userToLogin fetching
	userToLogin, err := userToLoginController.userToLoginService.GetUserToLogin(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, userToLogin)
}

func (userToLoginController *UserToLoginController) DeleteUserToLogin(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger userToLogin deletion
	if err := userToLoginController.userToLoginService.DeleteUserToLogin(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "UserToLogin deleted successfully",
	})
}

func (userToLoginController *UserToLoginController) ListUserToLogins(context *gin.Context) {
	// trigger all userToLogins fetching
	userToLogins, err := userToLoginController.userToLoginService.ListUserToLogins()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, userToLogins)
}

func (*UserToLoginController) PatchUserToLogin(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*UserToLoginController) OptionsUserToLogin(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*UserToLoginController) HeadUserToLogin(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
