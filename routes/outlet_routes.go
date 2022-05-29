package routes

import (
	"aulia-majoo-test/config"
	"aulia-majoo-test/internal/delivery"
	"aulia-majoo-test/internal/repository"
	"aulia-majoo-test/internal/usecase"

	"github.com/gin-gonic/gin"
)

func Outlet(route *gin.Engine) {
	db = config.AppConfig
	handlerUser := delivery.NewOutletHandler(usecase.NewOutletUsecase(repository.NewOutletRepository(db.DB)))

	v1 := route.Group("api/v1/")
	v1.GET("outlet/omzet/:id", authMiddleware(), handlerUser.FetchDailyIncomeOutletByID)

}
