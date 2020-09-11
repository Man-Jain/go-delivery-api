package services

import (
	"app/app/models"

	"golang.org/x/crypto/bcrypt"
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
func InsertUser(jsonData map[string]interface{}) (*models.User, error) {
	password := jsonData["password"].(string)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	obj := models.User{
		Profile: models.Profile{
			Email:    jsonData["email"].(string),
			Password: hashedPassword,
		},
	}
	if result := DB.Create(&obj); result.Error != nil {
		return &obj, result.Error
	}

	return &obj, nil
}

// InsertRootUser will insert a root user in db
func InsertRootUser(jsonData map[string]interface{}) (*models.User, error) {
	password := jsonData["password"].(string)
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	obj := models.User{
		Profile: models.Profile{
			Email:    jsonData["email"].(string),
			Password: hashedPassword,
		},
		Root: true,
	}
	if result := DB.Create(&obj); result.Error != nil {
		return &obj, result.Error
	}

	return &obj, nil
}
