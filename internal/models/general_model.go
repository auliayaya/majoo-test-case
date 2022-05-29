package models

type JSONResponse struct {
	Code   int         `json:"code"`
	Status interface{} `json:"status"`
	Data   interface{} `json:"data"`
	Meta   interface{} `json:"meta"`
}

type Responses struct {
	Code    int         `json:"code"`
	Status  interface{} `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type Meta struct {
	TotalPages      int  `json:"total_pages"`
	PageSize        int  `json:"page_size"`
	TotalItem       int  `json:"total_item"`
	HasNextPage     bool `json:"has_next_page"`
	HasPreviousPage bool `json:"has_previous_page"`
}
