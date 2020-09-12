package services

import (
	"app/app/models"
)

// QueryAllCookies return all users of application
func QueryAllCookies() (*[]models.Cookie, error) {
	cookies := []models.Cookie{}
	if result := DB.Find(&cookies); result.Error != nil {
		return &cookies, result.Error
	}
	return &cookies, nil
}

// QueryCookie return a single user object
func QueryCookie(id int) (*models.Cookie, error) {
	cookie := models.Cookie{}
	if result := DB.First(&cookie, id); result.Error != nil {
		return &cookie, result.Error
	}
	return &cookie, nil
}

// InsertCookie return a single user object
func InsertCookie(cookie models.Cookie) (*models.Cookie, error) {
	if result := DB.Create(&cookie); result.Error != nil {
		return &cookie, result.Error
	}
	return &cookie, nil
}
