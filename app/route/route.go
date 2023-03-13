package route

import (
	"airbnb/app/config"
	userData "airbnb/features/user/data"
	userHdl "airbnb/features/user/handler"
	userServ "airbnb/features/user/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userData := userData.New(db)
	userService := userServ.New(userData)
	userHandler := userHdl.New(userService)

	// AUTH
	e.POST("/login", userHandler.Login())

	// USER
	e.POST("/register", userHandler.Register())
	e.GET("/profile", userHandler.Profile(), middleware.JWT([]byte(config.JWTKey)))
	e.PUT("/profile", userHandler.Update(), middleware.JWT([]byte(config.JWTKey)))
	e.DELETE("/profile", userHandler.Delete(), middleware.JWT([]byte(config.JWTKey)))
}
