package model

type JsonResponse struct {
	RequestId string      `json:"request_id"`
	Status    int         `json:"status"`
	Messages  string      `json:"messages"`
	Data      interface{} `json:"data"`
}

type JsonResponseTotal struct {
	RequestId string      `json:"request_id"`
	Status    int         `json:"status"`
	Messages  string      `json:"messages"`
	Total     int         `json:"total"`
	Data      interface{} `json:"data"`
}

type Pagenation struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

type GormWhere struct {
	Where string
	Value []interface{}
}

type JsonResponsError struct {
	RequestId        string      `json:"request_id"`
	StatusCode       int         `json:"status_code"`
	ErrorCode        interface{} `json:"error_code"`
	ErrorMessage     interface{} `json:"error_message"`
	DeveloperMessage interface{} `json:"developer_message"`
}
