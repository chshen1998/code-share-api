package controllers

import (
	"gin_project/models"
	"gin_project/services"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SignUp(context *gin.Context) {
	var newUser models.User
	err := context.ShouldBindJSON(&newUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err = services.SignUp(&newUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Successfully signed up"})
}

func Login(context *gin.Context) {
	var user models.User
	context.ShouldBindJSON(&user)
	session := sessions.Default(context)

	userId, username, passwordDB := services.GetUser(user.Email)
	if userId == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Email does not have account"})
		return
	}

	if user.Password != passwordDB {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Email and password does not match"})
		return
	}

	session.Set("userId", userId)
	session.Set("username", username)
	if err := session.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
		return
	}
	context.JSON(http.StatusAccepted, models.User{userId, username, user.Email, user.Password})
}

func Logout(context *gin.Context) {
	session := sessions.Default(context)
	userId := session.Get("userId")
	if userId == nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session key"})
		return
	}
	session.Delete("userId")
	session.Delete("username")
	session.Save()
	context.JSON(http.StatusOK, gin.H{"message": "Sucessfully logged out"})
}

func GetSession(context *gin.Context) {
	session := sessions.Default(context)
	userId := session.Get("userId")
	if userId == nil {
		context.JSON(http.StatusOK, gin.H{"sessionFound": false})
		return
	}
	username := session.Get("username")
	context.JSON(http.StatusOK, gin.H{"sessionFound": true, "userId": userId, "username": username})
}

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	userId := session.Get("userId")
	if userId == nil {
		// Abort the request with the appropriate error code
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}
