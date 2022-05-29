package domain

import (
	"aulia-majoo-test/internal/dto"
	"aulia-majoo-test/internal/models"
	"time"

	"github.com/gin-gonic/gin"
)

type Merchant struct {
	ID           uint          `gorm:"primaryKey;column:id;"`
	UserID       uint          `gorm:"column:user_id"`
	MerchantName string        `gorm:"column:merchant_name"`
	Outlets      []Outlet      `json:"outlets" gorm:"foreignkey:MerchantID;references:ID"`
	Transactions []Transaction `json:"transactions"  gorm:"foreignkey:MerchantID;references:ID"`
	CreatedAt    time.Time     `gorm:"column:created_at"`
	CreatedBy    uint          `gorm:"column:created_by"`
	UpdatedAt    time.Time     `gorm:"column:updated_at"`
	UpdatedBy    uint          `gorm:"column:updated_by"`
}

func (Merchant) TableName() string {
	return "Merchants"
}

type (
	MerchantRepository interface {
		GetMerchantByUser(userID int64) (*Merchant, error)
		GetMerchantDailyIncome(id, start, end, offset, limit string) (resp []dto.MerchantDailyIncome, totalData int64, err error)
	}
	MerchantUsecase interface {
		DailyIncomeMerchantByIDQuery(userID string, id string, request dto.DatePaginationParameter) (models.JSONResponse, error)
	}
	MerchantDelivery interface {
		FetchDailyIncomeMerchantByID(*gin.Context)
	}
)
