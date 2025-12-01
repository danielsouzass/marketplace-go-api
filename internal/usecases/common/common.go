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

func BadRequestResponse(body types.Error) types.Response {
	return types.Response{
		Body: body,
		Code: http.StatusBadRequest,
	}
}

func CreatedResponse(body any) types.Response {
	return types.Response{
		Body: body,
		Code: http.StatusCreated,
	}
}
