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
func QueryUser(id uint) (*models.User, error) {
	user := models.User{}
	result := DB.First(&user, id)
	println(result.RowsAffected)
	return &user, nil
}

// InsertUser will insert a user in db
func InsertUser(obj models.User) (*models.User, error) {

	if result := DB.Create(&obj); result.Error != nil {
		return &obj, result.Error
	}

	return &obj, nil
}

// InsertRootUser will insert a root user in db
func InsertRootUser(obj models.User) (*models.User, error) {

	if result := DB.Create(&obj); result.Error != nil {
		return &obj, result.Error
	}

	return &obj, nil
}
