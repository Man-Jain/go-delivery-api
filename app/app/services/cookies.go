package services

import (
	"app/app/models"
)

// QueryAllCookies return all users of application
func QueryAllCookies() (*[]models.Cookie, error) {
	cookies := []models.Cookie{}
	result := DB.Find(&cookies)
	println(result)
	println(result.RowsAffected)
	return &cookies, nil
}

// QueryCookie return a single user object
func QueryCookie(id int) (*models.Cookie, error) {
	cookie := models.Cookie{}
	result := DB.First(&cookie, id)
	println(result.RowsAffected)
	return &cookie, nil
}

// InsertCookie return a single user object
func InsertCookie(jsonData map[string]interface{}) (*models.Cookie, error) {
	cookie := models.Cookie{
		Name:        jsonData["name"].(string),
		Description: jsonData["description"].(string),
		Price:       uint(jsonData["price"].(float64)),
		Quantity:    uint(jsonData["quantity"].(float64)),
	}

	result := DB.Create(&cookie)
	println(result.RowsAffected)
	return &cookie, nil
}
