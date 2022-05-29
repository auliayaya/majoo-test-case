package domain

import "time"

type Transaction struct {
	ID         uint      `gorm:"primaryKey;column:id;"`
	MerchantID uint      `gorm:"column:merchant_id"`
	OutletID   uint      `gorm:"column:outlet_id"`
	BillTotal  float64   `gorm:"column:bill_total"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	CreatedBy  uint      `gorm:"column:created_by"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
	UpdatedBy  uint      `gorm:"column:updated_by"`
}

func (Transaction) TableName() string {
	return "Transactions"
}
