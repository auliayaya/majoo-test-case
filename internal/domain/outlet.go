package domain

import (
	"aulia-majoo-test/internal/dto"
	"aulia-majoo-test/internal/models"
	"time"

	"github.com/gin-gonic/gin"
)

type Outlet struct {
	ID           uint          `gorm:"primaryKey;column:id;"`
	MerchantID   uint          `gorm:"column:merchant_id"`
	OutletName   string        `gorm:"column:outlet_name"`
	Transactions []Transaction `json:"transactions"  gorm:"foreignkey:OutletID;references:ID"`
	CreatedAt    time.Time     `gorm:"column:created_at"`
	CreatedBy    uint          `gorm:"column:created_by"`
	UpdatedAt    time.Time     `gorm:"column:updated_at"`
	UpdatedBy    uint          `gorm:"column:updated_by"`
}

func (Outlet) TableName() string {
	return "Outlets"
}

type (
	OutletRepository interface {
		GetOutletByIDAndUser(userID, id int64) (*Outlet, error)
		GetOutletDailyIncome(id, start, end, offset, limit string) (resp []dto.OutletDailyIncome, totalData int64, err error)
	}
	OutletUsecase interface {
		DailyIncomeOutletByIDQuery(userID string, id string, request dto.DatePaginationParameter) (models.JSONResponse, error)
	}
	OutletDelivery interface {
		FetchDailyIncomeOutletByID(*gin.Context)
	}
)
