package api

import (
	"marketplace/internal/jsonutils"
	"marketplace/internal/types"
	"marketplace/internal/usecases/common"
	"net/http"
)

func (api *API) handleGetPaymentMethods(w http.ResponseWriter, r *http.Request) {
	response, err := api.PaymentMethodsService.GetPaymentMethods(r.Context())
	if err != nil {
		jsonutils.SendJSON(w, common.BadRequestResponse(types.Error{
			Message: err.Error(),
		}))
		return
	}

	jsonutils.SendJSON(w, common.OKResponse(response))
}
