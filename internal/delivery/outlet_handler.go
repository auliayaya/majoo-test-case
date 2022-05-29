package delivery

import (
	"aulia-majoo-test/internal/domain"
	"aulia-majoo-test/internal/dto"
	"aulia-majoo-test/pkg/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type outletHandler struct {
	ucase domain.OutletUsecase
}

func NewOutletHandler(ucase domain.OutletUsecase) domain.OutletDelivery {
	return &outletHandler{
		ucase: ucase,
	}
}

// Outlet Daily Income By ID  godoc
// @Summary Outlet Daily Income
// @Description Outlet Daily Income
// @Tags Users
// @Accept json
// @Produce json
// @Param User body dto.OutletRequestDTO true "Outlet"
// @Success 200 {object} dto.OutletResponseDTO
// @Router /outlet/income/:id [get]
func (c *outletHandler) FetchDailyIncomeOutletByID(ctx *gin.Context) {
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

	parameter := ctx.Param("id")
	// req.StartDate = "2021-11-01"
	// req.EndDate = "2021-11-30"

	res, err := c.ucase.DailyIncomeOutletByIDQuery(id, parameter, req)

	if err != nil {
		ctx.JSON(res.Code, err)
	}
	ctx.JSON(res.Code, res)
}
