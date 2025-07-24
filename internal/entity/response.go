package entity

type SuccessResponse struct {
	Message    string      `json:"message"`
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
	DateTime   string      `json:"dateTime"`
}

type ErrorResponse struct {
	Message    string      `json:"message"`
	StatusCode int         `json:"statusCode"`
	DateTime   string      `json:"dateTime"`
	Data       interface{} `json:"data"`
}

type PaginationResponse struct {
	TotalItems  int         `json:"totalItems"`
	TotalPages  int         `json:"totalPages"`
	CurrentPage int         `json:"currentPage"`
	Items       interface{} `json:"items"`
}
