package src

import (
	"api-2/src/delivery"
	"api-2/src/repository"
	"api-2/src/usecase"
	"api-2/utils/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type server struct {
	httpServer *gin.Engine
	database   *gorm.DB
}

type Server interface {
	Run()
}

func InitServer() *server {
	e := gin.Default()

	s := &server{
		httpServer: e,
		database:   db.NewDBGorm().Master(),
	}

	return s
}

func (s *server) Run() {
	s.httpServer.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test request success",
		})
	})

	userRepository := repository.NewRepository(s.database)
	userUsecase := usecase.NewUserUsecase(userRepository)
	authDelivery := delivery.NewDelivery(userUsecase)

	authGroup := s.httpServer.Group("auth")
	authDelivery.Mount(authGroup)

	s.httpServer.Run(":8082")
}
