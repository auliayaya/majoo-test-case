package repository

import (
	"aulia-majoo-test/internal/domain"
	"aulia-majoo-test/internal/dto"
	"errors"

	"gorm.io/gorm"
)

type outletRepository struct {
	DB *gorm.DB
}

func NewOutletRepository(DB *gorm.DB) domain.OutletRepository {
	return &outletRepository{
		DB: DB,
	}
}

func (r *outletRepository) GetOutletDailyIncome(id, start, end, offset, limit string) (resp []dto.OutletDailyIncome, totalData int64, err error) {

	err = r.DB.Raw(getOutletDailyIncome, start, id, end, limit, offset).Scan(&resp).Error

	if err != nil {
		return resp, 0, err
	}
	var ttlData []dto.OutletDailyIncome
	err = r.DB.Raw(getOutletDailyIncomeNoLimitOffset, start, id, end).Scan(&ttlData).Error
	if err != nil {
		return resp, 0, err
	}
	return resp, int64(len(ttlData)), nil
}

func (r *outletRepository) GetOutletByIDAndUser(userID, id int64) (*domain.Outlet, error) {
	var outletEntity domain.Outlet

	err := r.DB.Raw(getOutletByIDAndUserID, id, userID).First(&outletEntity).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &outletEntity, err
}
