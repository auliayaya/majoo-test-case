package usecase

import (
	"aulia-majoo-test/internal/domain"
	"aulia-majoo-test/internal/dto"
	"aulia-majoo-test/internal/models"
	"aulia-majoo-test/pkg/c"
	"strconv"
	"time"
)

type outletUsecase struct {
	uRepo domain.OutletRepository
}

func NewOutletUsecase(uRepo domain.OutletRepository) domain.OutletUsecase {
	return &outletUsecase{uRepo: uRepo}
}

func (u *outletUsecase) DailyIncomeOutletByIDQuery(userID string, id string, request dto.DatePaginationParameter) (models.JSONResponse, error) {

	_, err := time.Parse(c.DateFormat, request.StartDate)
	if err != nil {
		return models.BadRequestPaginationResponses("Start date parameter format should be YYYY-MM-DD"), nil
	}
	_, err = time.Parse(c.DateFormat, request.EndDate)
	if err != nil {
		return models.BadRequestPaginationResponses("End date parameter format should be YYYY-MM-DD"), nil
	}

	if request.PageNumber == 0 {
		request.PageNumber = c.DefaultPageNumber
	}
	if request.PageSize == 0 {
		request.PageSize = c.DefaultPageSize
	}

	offset := request.PageSize * (request.PageNumber - 1)

	user, _ := strconv.Atoi(userID)
	ids, _ := strconv.Atoi(id)

	outlet, err := u.uRepo.GetOutletByIDAndUser(int64(user), int64(ids))
	if err != nil {
		return models.InternalServerPaginationResponses("Internal Server Error"), err
	}
	if outlet == nil {
		return models.NotFoundPaginationResponses("User didnt has merchant", nil), nil
	}

	outletID := strconv.FormatInt(int64(outlet.ID), 10)

	if outletID != id {
		return models.ForbiddenExceptionPagination(), nil
	}

	omzet, totalData, err := u.uRepo.GetOutletDailyIncome(outletID, request.StartDate, request.EndDate, strconv.Itoa(offset), strconv.Itoa(request.PageSize))
	if err != nil {
		return models.InternalServerPaginationResponses("Internal Server Error"), err
	}

	c := float64(totalData) / float64(request.PageSize)
	totalPages := int(c)
	if totalPages == 0 {
		totalPages = 1
	}

	conditionNextPage := false
	if request.PageNumber < totalPages {
		conditionNextPage = true
	}

	meta := models.Meta{
		PageSize:        request.PageSize,
		TotalPages:      totalPages,
		TotalItem:       int(totalData),
		HasNextPage:     conditionNextPage,
		HasPreviousPage: !(request.PageNumber <= 1),
	}

	return models.SuccessPaginationResponses(omzet, meta), nil
}
