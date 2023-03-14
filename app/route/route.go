package route

import (
	"airbnb/app/config"
	userData "airbnb/features/user/data"
	userHdl "airbnb/features/user/handler"
	userServ "airbnb/features/user/services"

	homestayData "airbnb/features/homestay/data"
	homestayHdl "airbnb/features/homestay/handler"
	homestaySrv "airbnb/features/homestay/services"

	feedbackData "airbnb/features/feedback/data"
	feedbackHdl "airbnb/features/feedback/handler"
	feedbackSrv "airbnb/features/feedback/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userData := userData.New(db)
	userService := userServ.New(userData)
	userHandler := userHdl.New(userService)

	homestayData := homestayData.New(db)
	homestayService := homestaySrv.New(homestayData)
	homestayHandler := homestayHdl.New(homestayService)

	feedbackData := feedbackData.New(db)
	feedbackService := feedbackSrv.New(feedbackData)
	feedbackHandler := feedbackHdl.New(feedbackService)

	// AUTH
	e.POST("/login", userHandler.Login())

	// USER
	e.POST("/register", userHandler.Register())
	e.GET("/profile", userHandler.Profile(), middleware.JWT([]byte(config.JWTKey)))
	e.PUT("/profile", userHandler.Update(), middleware.JWT([]byte(config.JWTKey)))
	e.DELETE("/profile", userHandler.Delete(), middleware.JWT([]byte(config.JWTKey)))

	// HOMESTAY
	e.POST("/homestays", homestayHandler.Add(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/homestays", homestayHandler.ShowAll(), middleware.JWT([]byte(config.JWTKey)))
	e.GET("/homestays/:homestay_id", homestayHandler.ShowDetail(), middleware.JWT([]byte(config.JWTKey)))
	e.PUT("/homestays/:homestay_id", homestayHandler.Update(), middleware.JWT([]byte(config.JWTKey)))
	e.DELETE("/homestays/:homestay_id", homestayHandler.Delete(), middleware.JWT([]byte(config.JWTKey)))

	// FEEDBACK
	e.POST("/feedbacks", feedbackHandler.AddFeedback(), middleware.JWT([]byte(config.JWTKey)))
}
