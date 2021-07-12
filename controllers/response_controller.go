package controllers

import (
	"gin_project/models"
	"gin_project/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ResponseController contains the methods used to read the request data, which is then passed into the ResponseController.
type ResponseController interface {
	GetLikes(context *gin.Context) int
	AddLike(context *gin.Context) int
	GetComments(context *gin.Context) []models.Comment
	AddComment(context *gin.Context) []models.Comment
}

type ResponseControllerImpl struct {
	service services.ResponseService
}

func NewResponseController(service services.ResponseService) ResponseController {
	return &ResponseControllerImpl{
		service: service,
	}
}

// Get the likes count of the post with id of postId
func (controller *ResponseControllerImpl) GetLikes(context *gin.Context) int {
	postId, err := strconv.Atoi(context.Query("postId"))
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	return controller.service.GetLikes(postId)
}

// Adds a like from user to post
func (controller *ResponseControllerImpl) AddLike(context *gin.Context) int {
	var like models.Like
	err := context.ShouldBindJSON(&like)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	err = controller.service.AddLike(like)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	return controller.service.GetLikes(like.PostId)
}

// Gets all commments related to post with postId
func (controller *ResponseControllerImpl) GetComments(context *gin.Context) []models.Comment {
	postId, err := strconv.Atoi(context.Query("postId"))
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	return controller.service.GetComments(postId)
}

// Add comment from user to post
func (controller *ResponseControllerImpl) AddComment(context *gin.Context) []models.Comment {
	var comment models.Comment
	err := context.ShouldBindJSON(&comment)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	err = controller.service.AddComment(&comment)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	return controller.service.GetComments(comment.PostId)
}
