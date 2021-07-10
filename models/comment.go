package models

type Comment struct {
	Id       int    `json:"id"`
	Username string `json:"username" binding:"required"`
	PostId   int    `json:"postId" binding:"required"`
	Comment  string `json:"comment" binding:"required"`
}
