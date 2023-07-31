package delivery

import "github.com/labstack/echo/v4"

func (d *delivery) Auth(e *echo.Group) {
	e.POST("/register", d.RegisterHandler)
	e.POST("/login", d.Login)

	e.POST("/logout", d.Logout)
}

func (d *delivery) User(e *echo.Group) {
	e.GET("", d.GetUserByID)
}
