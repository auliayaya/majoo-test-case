package routes

import (
	"aulia-majoo-test/config"
	"aulia-majoo-test/internal/delivery"
	"aulia-majoo-test/internal/repository"
	"aulia-majoo-test/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var db *config.APP

func User(route *gin.Engine) {
	db = config.AppConfig
	handlerUser := delivery.NewUserHandler(usecase.NewUserUsecase(repository.NewUserRepository(db.DB)))
	route.GET("/api/v1/checkhealth", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "success", "status": true})
	})
	v1 := route.Group("api/v1/")
	v1.POST("auth/login", handlerUser.Login)
	v1.POST("auth/register", handlerUser.Register)
	docs := route.Group("")
	docs.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
