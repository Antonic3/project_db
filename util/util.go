package util

import "project-db/model"

func CreateResponse(isSuccess bool, data any, errorMessage string) model.Response {
	return model.Response{
		Success: isSuccess,
		Data:    data,
		Message: errorMessage,
	}
}
