package database

import (
	"day_2/config"
	"day_2/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserById(id int) (interface{}, error) {
	var user models.User
	if err := config.DB.Model(models.User{}).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func DeleteUserById(id int) (interface{}, error) {
	if err := config.DB.Delete(&models.User{}, id).Error; err != nil {
		return nil, err
	}
	return nil, nil
}

func CreateUser(user models.User) (interface{}, error) {
	err := config.DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func UpdateUser(id int, user models.User) (interface{}, error) {
	err := config.DB.Where("id = ?", id).Updates(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func LoginUser(user models.User) (models.User, error) {
	err := config.DB.Model(models.User{}).
		Where("email = ? AND password = ?", user.Email, user.Password).
		First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
