package models

type Post struct {
	Id          int    `json:"id"`
	Title       string `json:"title" binding"required"`
	Description string `json:"description"`
	Code        string `json:"code" binding:"required"`
	AuthorName  string `json:"authorName" binding:"required"`
	Language    string `json:"language" binding:"required"`
	LikesCount  int    `json:"likesCount"`
}
