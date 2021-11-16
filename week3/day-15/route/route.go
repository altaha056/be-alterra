package route

import (
	"project/constants"
	c "project/controller"
	m "project/middleware"

	"github.com/labstack/echo"
	em "github.com/labstack/echo/middleware"
)

func New() *echo.Echo{
	e:=echo.New()
	e.GET("/users", c.GetUserController)
	m.LogMiddleware(e)
	e.POST("/users",c.CreateUserController)
	e.POST("/login",c.LoginUserController)
	
	eAuth:=e.Group("/auth")
	eAuth.Use(em.BasicAuth(m.BasicAuthDB))
	eAuth.GET("/users",c.GetUserController)

	eJwt:=e.Group("/jwt")
	eJwt.Use(em.JWT([]byte(constants.SECRET_JWT)))
	eJwt.GET("/users",c.GetUserController)
	
	
	return e
}