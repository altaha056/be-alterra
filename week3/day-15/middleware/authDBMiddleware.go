package middleware

import (
	"project/config"
	"project/model"

	"github.com/labstack/echo"
)

func BasicAuthDB(username,password string, c echo.Context)(bool,error){
	var db = config.DB
	var user model.Users
	err:=db.Where("email=? AND password =?",username,password).First(&user).Error;
	if err!=nil{
		return false,nil
	}
	return true,nil
}