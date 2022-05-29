package repository

import (
	"aulia-majoo-test/internal/domain"
	"aulia-majoo-test/internal/dto"
	"errors"

	"gorm.io/gorm"
)

type merchantRepository struct {
	DB *gorm.DB
}

func NewMerchantRepository(DB *gorm.DB) domain.MerchantRepository {
	return merchantRepository{
		DB: DB,
	}
}

func (r merchantRepository) CreateMerchant(merchant domain.Merchant) (domain.Merchant, error) {
	err := r.DB.Debug().Create(&merchant).Error
	if err != nil {
		return domain.Merchant{}, err
	}

	return merchant, nil
}

func (r merchantRepository) GetMerchantByUser(userID int64) (*domain.Merchant, error) {
	var merchantEntity domain.Merchant
	err := r.DB.Preload("Outlets").
		Preload("Transactions").Where("user_id = ?", userID).First(&merchantEntity).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &merchantEntity, err
}

func (r merchantRepository) GetMerchantDailyIncome(id, start, end, offset, limit string) (resp []dto.MerchantDailyIncome, totalData int64, err error) {

	err = r.DB.Raw(getMerchantDailyIncome, start, id, end, limit, offset).Scan(&resp).Error

	if err != nil {
		return resp, 0, err
	}
	var ttlData []dto.MerchantDailyIncome
	err = r.DB.Raw(getMerchanDailyIncomeNoLimitOffset, start, id, end).Scan(&ttlData).Error
	if err != nil {
		return resp, 0, err
	}
	return resp, int64(len(ttlData)), nil
}
