package models

type LoginResponse struct {
	Token        Token  `json:"token"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	MerchantName string `json:"merchant_name"`
}

type RegisterResponse struct {
	Token        Token  `json:"token"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	MerchantName string `json:"merchant_name"`
}

type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}
