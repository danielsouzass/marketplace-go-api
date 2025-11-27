package api

import (
	"marketplace/internal/jsonutils"
	"marketplace/internal/types"
	"marketplace/internal/usecases/common"
	"marketplace/internal/usecases/user"
	"net/http"
)

func (api *API) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	data, err := jsonutils.ReadJSON[user.CreateUserRequest](r)
	if err != nil {
		jsonutils.SendJSON(w, common.InvalidJSONResponse(err))
		return
	}

	if err := api.Validator.Struct(data); err != nil {
		jsonutils.SendJSON(w, common.ValidationErrorResponse(err))
		return
	}

	id, err := api.UserService.CreateUser(r.Context(), data)
	if err != nil {
		jsonutils.SendJSON(w, user.CreateUser400Response(types.Error{
			Message: err.Error(),
		}))
		return
	}

	jsonutils.SendJSON(w, user.CreateUser201Response(user.CreateUserResponse{
		Id: id.String(),
	}))
}

func (api *API) handleGetUser(w http.ResponseWriter, r *http.Request) {
	response, err := api.UserService.GetUser(r.Context())
	if err != nil {
		jsonutils.SendJSON(w, user.GetUser404Response(types.Error{
			Message: err.Error(),
		}))
		return
	}

	jsonutils.SendJSON(w, user.GetUser200Response(response))
}
