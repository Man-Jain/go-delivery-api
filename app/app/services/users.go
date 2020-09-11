package services

import (
	"app/app/models"
)

// QueryAllUsers return all users of application
func QueryAllUsers() (*[]models.User, error) {
	users := []models.User{}
	result := DB.Find(&users)
	println("the users are", result)
	println(result.RowsAffected)
	return &users, nil
}

// QueryUser return a single user object
func QueryUser(id int) (*models.User, error) {
	user := models.User{}
	result := DB.First(&user, id)
	println(result.RowsAffected)
	return &user, nil
}
