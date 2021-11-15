package controller

import (
	"net/http"
	"project/config"
	"project/model"

	"github.com/labstack/echo"
)

func GetUserController(c echo.Context) error {
	var users []model.Users
	err := config.DB.Find(&users).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
	})
}

func CreateUserController(c echo.Context) error {

	user := model.Users{}
	c.Bind(&user)

	err := config.DB.Save(&user).Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})

}