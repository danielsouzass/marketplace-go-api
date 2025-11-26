package common

import (
	"marketplace/internal/types"
	"net/http"
)

func InvalidJSONResponse(err error) types.Response {
	return types.Response{
		Body: types.Error{
			Message: "invalid JSON: " + err.Error(),
		},
		Code: http.StatusBadRequest,
	}
}

func ValidationErrorResponse(err error) types.Response {
	return types.Response{
		Body: types.Error{
			Message: "validation error: " + err.Error(),
		},
		Code: http.StatusUnprocessableEntity,
	}
}
