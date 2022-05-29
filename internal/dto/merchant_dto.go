package dto

type (
	MerchantDailyIncome struct {
		MerchantID   int    `json:"merchant_id"`
		MerchantName string `json:"merchant_name"`
		Date         string `json:"date"`
		Total        int    `json:"total"`
	}
)
type (
	DateParameter struct {
		StartDate string `query:"start_date"`
		EndDate   string `query:"end_date"`
	}

	DatePaginationParameter struct {
		StartDate  string `query:"start_date"`
		EndDate    string `query:"end_date"`
		PageSize   int    `query:"page_number"`
		PageNumber int    `query:"page_number"`
	}
)
