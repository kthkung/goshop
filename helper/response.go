package helper

import "strings"

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"errors"`
	Page    int         `json:"page"`
	Limit   int         `json:"limit"`
	Total   int64       `json:"total"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(status bool, message string, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Error:   nil,
		Data:    data,
	}
	return res
}

func ListResponse(status bool, message string, page int, limit int, total int64, data interface{}) Response {
	res := Response{
		Status:  status,
		Message: message,
		Error:   nil,
		Page:    page,
		Limit:   limit,
		Total:   total,
		Data:    data,
	}
	return res
}

func BuildErrorResponse(message string, err string, data interface{}) Response {
	splittedError := strings.Split(err, "\n")
	res := Response{
		Status:  false,
		Message: message,
		Error:   splittedError,
		Data:    data,
	}
	return res
}
