package dto

type LoginRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequestDTO struct {
	Name         string `json:"name"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	MerchantName string `json:"merchant_name"`
}
