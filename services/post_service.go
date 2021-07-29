package services

import (
	"fmt"
	"gin_project/config"
	"gin_project/models"
)

type PostService interface {
	Save(*models.Post) error
	FindAll() []models.Post
	FindFiltered(string, string) []models.Post
	Get(string) []models.Post
	Delete(int) error
}

type PostServiceImpl struct{}

func New() PostService {
	return &PostServiceImpl{}
}

func (service *PostServiceImpl) Save(post *models.Post) error {
	sqlStatement := "INSERT INTO Posts (title, description, code, authorName, language) VALUES ($1, $2, $3, $4, $5)"
	_, err := config.DB.Exec(sqlStatement, post.Title, post.Description, post.Code, post.AuthorName, post.Language)
	return err
}

func (service *PostServiceImpl) FindAll() []models.Post {
	var title, description, code, authorName, language string
	var id, likesCount int
	var posts []models.Post
	rows, _ := config.DB.Query(
		"SELECT P.id, P.title, P.description, P.code, P.authorName, P.language, count(L.username) as likesCount " +
			"FROM Posts P left join Likes L ON P.id = L.postId " +
			"GROUP BY P.id ORDER BY likesCount desc")
	for rows.Next() {
		err := rows.Scan(&id, &title, &description, &code, &authorName, &language, &likesCount)
		if err != nil {
			panic(err)
		}
		posts = append(posts, models.Post{id, title, description, code, authorName, language, likesCount})
	}
	return posts
}

func (service *PostServiceImpl) Get(username string) []models.Post {
	var title, description, code, authorName, language string
	var id, likesCount int
	var posts []models.Post
	sqlStatement := fmt.Sprintf("SELECT * FROM Posts WHERE authorname = '%s'", username)
	rows, _ := config.DB.Query(sqlStatement)
	for rows.Next() {
		err := rows.Scan(&id, &title, &description, &code, &authorName, &language, &likesCount)
		if err != nil {
			panic(err)
		}
		posts = append(posts, models.Post{id, title, description, code, authorName, language, likesCount})
	}
	return posts
}

func (service *PostServiceImpl) Delete(postId int) error {
	sqlStatement := fmt.Sprintf("DELETE FROM Posts WHERE id = %d", postId)
	_, err := config.DB.Exec(sqlStatement)
	return err
}

func (service *PostServiceImpl) FindFiltered(keywords string, languagef string) []models.Post {
	var title, description, code, authorName, language string
	var id int
	var posts []models.Post
	var sqlStatement string

	if languagef == "" {
		sqlStatement = fmt.Sprintf("SELECT * FROM Posts WHERE (title ~ '%s' OR description ~ '%s')", keywords, keywords)
	} else {
		sqlStatement = fmt.Sprintf("SELECT * FROM Posts WHERE language = '%s' AND (title ~ '%s' OR description ~ '%s')", languagef, keywords, keywords)
	}
	rows, _ := config.DB.Query(sqlStatement)
	for rows.Next() {
		err := rows.Scan(&id, &title, &description, &code, &authorName, &language)
		if err != nil {
			panic(err)
		}
		posts = append(posts, models.Post{id, title, description, code, authorName, language, 2})
	}
	return posts
}
