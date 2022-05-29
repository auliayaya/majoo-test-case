package routes

import (
	"aulia-majoo-test/config"
	"aulia-majoo-test/internal/delivery"
	"aulia-majoo-test/internal/repository"
	"aulia-majoo-test/internal/usecase"

	"github.com/gin-gonic/gin"
)

func Merchant(route *gin.Engine) {
	db = config.AppConfig
	handlerUser := delivery.NewMerchantHandler(usecase.NewMerchantUsecase(repository.NewMerchantRepository(db.DB)))

	v1 := route.Group("api/v1/")
	v1.GET("merchant/omzet/:id", authMiddleware(), handlerUser.FetchDailyIncomeMerchantByID)

}
