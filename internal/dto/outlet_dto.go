package dto

type (
	OutletDailyIncome struct {
		MerchantID   int    `json:"merchant_id"`
		MerchantName string `json:"merchant_name"`
		Date         string `json:"date"`
		Total        int    `json:"total"`
	}
)
