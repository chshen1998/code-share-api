package services

import (
	"fmt"
	"gin_project/config"
	"gin_project/models"
)

func SignUp(user *models.User) error {
	sqlStatement := `INSERT INTO Users (username, email, password) VALUES ($1, $2, $3)`
	_, err := config.DB.Exec(sqlStatement, user.Username, user.Email, user.Password)
	return err
}

func GetUser(email string) (int, string, string) {
	var username, password string
	var id int
	sqlStatement := fmt.Sprintf("SELECT id, username, password FROM Users WHERE email = '%s' LIMIT 1", email)
	err := config.DB.QueryRow(sqlStatement).Scan(&id, &username, &password)
	if err != nil {
		fmt.Println(err)
	}
	return id, username, password
}
