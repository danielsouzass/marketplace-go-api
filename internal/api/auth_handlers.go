package api

import (
	"marketplace/internal/jsonutils"
	"marketplace/internal/types"
	"marketplace/internal/usecases/common"
	"marketplace/internal/usecases/user"
	"net/http"
)

func (api *API) handleLoginUser(w http.ResponseWriter, r *http.Request) {
	data, err := jsonutils.ReadJSON[user.LoginUserRequest](r)
	if err != nil {
		jsonutils.SendJSON(w, common.InvalidJSONResponse(err))
		return
	}

	if err := api.Validator.Struct(data); err != nil {
		jsonutils.SendJSON(w, common.ValidationErrorResponse(err))
		return
	}

	tokens, err := api.UserService.AuthenticateUser(r.Context(), data)
	if err != nil {
		jsonutils.SendJSON(w, user.LoginUser400Response(types.Error{
			Message: err.Error(),
		}))
		return
	}

	jsonutils.SendJSON(w, user.LoginUser201Response(tokens))
}

func (api *API) handleUserRefreshToken(w http.ResponseWriter, r *http.Request) {
	data, err := jsonutils.ReadJSON[user.RefreshTokenUserRequest](r)
	if err != nil {
		jsonutils.SendJSON(w, common.InvalidJSONResponse(err))
		return
	}

	if err := api.Validator.Struct(data); err != nil {
		jsonutils.SendJSON(w, common.ValidationErrorResponse(err))
		return
	}

	refreshToken, err := api.UserService.RefreshUserToken(r.Context(), data.RefreshToken)
	if err != nil {
		jsonutils.SendJSON(w, user.RefreshTokenUser400Response(types.Error{
			Message: err.Error(),
		}))
		return
	}

	jsonutils.SendJSON(w, user.RefreshTokenUser200Response(refreshToken))
}
