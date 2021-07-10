package services

import (
	"fmt"
	"gin_project/config"
	"gin_project/models"
)

type ResponseService interface {
	GetLikes(postId int) int
	AddLike(like models.Like) error
	GetComments(postId int) []models.Comment
	AddComment(comment *models.Comment) error
}

type ResponseServiceImpl struct{}

func NewResponseService() ResponseService {
	return &ResponseServiceImpl{}
}

func (service *ResponseServiceImpl) GetLikes(postId int) int {
	var likeCount int
	sqlStatement := fmt.Sprintf("select count(*) as likesCount from Posts P left join Likes L on (P.id = L.postId) where P.id = %d", postId)
	err := config.DB.QueryRow(sqlStatement).Scan(&likeCount)
	if err != nil {
		fmt.Println(err)
	}
	return likeCount
}

func (service *ResponseServiceImpl) AddLike(like models.Like) error {
	sqlStatement := fmt.Sprintf("INSERT INTO LIKES (username, postId) VALUES ('%s', %d)", like.Username, like.PostId)
	_, err := config.DB.Exec(sqlStatement)
	return err
}

func (service *ResponseServiceImpl) GetComments(postIdf int) []models.Comment {
	var username, comment string
	var id, postId int
	var comments []models.Comment
	sqlStatement := fmt.Sprintf("SELECT * FROM Comments WHERE postId = %d", postIdf)
	rows, _ := config.DB.Query(sqlStatement)
	for rows.Next() {
		err := rows.Scan(&id, &username, &postId, &comment)
		if err != nil {
			panic(err)
		}
		comments = append(comments, models.Comment{id, username, postId, comment})
	}
	return comments
}

func (service *ResponseServiceImpl) AddComment(comment *models.Comment) error {
	sqlStatement := fmt.Sprintf("INSERT INTO Comments (username, postId, comment) VALUES ('%s', %d, '%s')", comment.Username, comment.PostId, comment.Comment)
	_, err := config.DB.Exec(sqlStatement)
	return err
}
