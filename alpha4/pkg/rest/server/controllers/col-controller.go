package controllers

import (
	"github.com/azar-intelops/microservice-demo/alpha4/pkg/rest/server/models"
	"github.com/azar-intelops/microservice-demo/alpha4/pkg/rest/server/services"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type ColController struct {
	colService *services.ColService
}

func NewColController() (*ColController, error) {
	colService, err := services.NewColService()
	if err != nil {
		return nil, err
	}
	return &ColController{
		colService: colService,
	}, nil
}

func (colController *ColController) CreateCol(context *gin.Context) {
	// validate input
	var input models.Col
	if err := context.ShouldBindJSON(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// trigger col creation
	if _, err := colController.colService.CreateCol(&input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Col created successfully"})
}

func (colController *ColController) UpdateCol(context *gin.Context) {
	// validate input
	var input models.Col
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

	// trigger col update
	if _, err := colController.colService.UpdateCol(id, &input); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Col updated successfully"})
}

func (colController *ColController) FetchCol(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger col fetching
	col, err := colController.colService.GetCol(id)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, col)
}

func (colController *ColController) DeleteCol(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// trigger col deletion
	if err := colController.colService.DeleteCol(id); err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Col deleted successfully",
	})
}

func (colController *ColController) ListCols(context *gin.Context) {
	// trigger all cols fetching
	cols, err := colController.colService.ListCols()
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, cols)
}

func (*ColController) PatchCol(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "PATCH",
	})
}

func (*ColController) OptionsCol(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "OPTIONS",
	})
}

func (*ColController) HeadCol(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "HEAD",
	})
}
