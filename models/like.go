package models

type Like struct {
	Id       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	PostId   int    `json:"postId" binding:"required"`
}
