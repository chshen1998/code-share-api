package controllers

import (
	"gin_project/models"
	"gin_project/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PostController contains the methods used to read the request data, which is then passed into the PostService.
type PostController interface {
	FindAll() []models.Post
	FindFiltered(context *gin.Context) []models.Post
	Save(context *gin.Context) error
	Get(context *gin.Context) []models.Post
	Delete(context *gin.Context) error
}

type PostControllerImpl struct {
	service services.PostService
}

func New(service services.PostService) PostController {
	return &PostControllerImpl{
		service: service,
	}
}

// Gets all posts in database
func (controller *PostControllerImpl) FindAll() []models.Post {
	return controller.service.FindAll()
}

// Adds a new post to the database
func (controller *PostControllerImpl) Save(context *gin.Context) error {
	var post models.Post
	err := context.ShouldBindJSON(&post)
	if err != nil {
		return err
	}
	err = controller.service.Save(&post)
	return err
}

// Gets all posts posted by the requesting user
func (controller *PostControllerImpl) Get(context *gin.Context) []models.Post {
	username := context.Query("username")
	return controller.service.Get(username)
}

// Deletes a post
func (controller *PostControllerImpl) Delete(context *gin.Context) error {
	postId, err := strconv.Atoi(context.Query("postId"))
	if err != nil {
		return err
	}
	return controller.service.Delete(postId)
}

// Finds all posts that are relevant to the keywords and language filter
func (controller *PostControllerImpl) FindFiltered(context *gin.Context) []models.Post {
	keywords := context.Query("keywords")
	language := context.Query("language")
	return controller.service.FindFiltered(keywords, language)
}
