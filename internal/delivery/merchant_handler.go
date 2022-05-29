package delivery

import (
	"aulia-majoo-test/internal/domain"
	"aulia-majoo-test/internal/dto"
	"aulia-majoo-test/pkg/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type merchantHandler struct {
	ucase domain.MerchantUsecase
}

func NewMerchantHandler(ucase domain.MerchantUsecase) domain.MerchantDelivery {
	return &merchantHandler{
		ucase: ucase,
	}
}

// Merchant Daily Income By ID  godoc
// @Summary Merchant Daily Income
// @Description Merchant Daily Income
// @Tags Users
// @Accept json
// @Produce json
// @Param User body dto.MerchantRequestDTO true "Merchant"
// @Success 200 {object} dto.MerchantResponseDTO
// @Router /merchant/income/:id [get]
func (c *merchantHandler) FetchDailyIncomeMerchantByID(ctx *gin.Context) {
	start_date := ctx.Query("start_date")
	end_date := ctx.Query("end_date")
	page_size := ctx.Query("page_size")
	page_number := ctx.Query("page_number")
	pSize, _ := strconv.Atoi(page_size)
	pNumber, _ := strconv.Atoi(page_number)

	req := dto.DatePaginationParameter{
		StartDate:  start_date,
		EndDate:    end_date,
		PageSize:   pSize,
		PageNumber: pNumber,
	}
	// get payload auth
	id, err := util.AuthContextToID(ctx)
	if err != nil {
		ctx.JSON(http.StatusForbidden, err)
		return
	}

	// err := models.ValidateRequest(ctx, &request)
	// if err != nil {
	// 	return err
	// }

	parameter := ctx.Param("id")
	// req.StartDate = "2021-11-01"
	// req.EndDate = "2021-11-30"

	res, err := c.ucase.DailyIncomeMerchantByIDQuery(id, parameter, req)

	if err != nil {
		ctx.JSON(res.Code, err)
	}
	ctx.JSON(res.Code, res)
}
