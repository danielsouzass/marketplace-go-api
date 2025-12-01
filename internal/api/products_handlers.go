package api

import (
	"marketplace/internal/jsonutils"
	"marketplace/internal/types"
	"marketplace/internal/usecases/common"
	"marketplace/internal/usecases/product"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (api *API) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	data, err := jsonutils.ReadJSON[product.CreateProductRequest](r)
	if err != nil {
		jsonutils.SendJSON(w, common.InvalidJSONResponse(err))
		return
	}

	if err := api.Validator.Struct(data); err != nil {
		jsonutils.SendJSON(w, common.ValidationErrorResponse(err))
		return
	}

	productID, err := api.ProductService.CreateProduct(r.Context(), data)
	if err != nil {
		jsonutils.SendJSON(w, common.BadRequestResponse(types.Error{
			Message: err.Error(),
		}))
		return
	}

	jsonutils.SendJSON(w, common.CreatedResponse(product.CreateProductResponse{
		Id: productID.String(),
	}))
}

func (api *API) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := api.ProductService.GetProducts(r.Context())
	if err != nil {
		jsonutils.SendJSON(w, common.BadRequestResponse(types.Error{
			Message: err.Error(),
		}))
		return
	}

	jsonutils.SendJSON(w, common.OKResponse(products))
}

func (api *API) handleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	err := api.ProductService.DeleteProduct(r.Context(), chi.URLParam(r, "id"))
	if err != nil {
		jsonutils.SendJSON(w, common.BadRequestResponse(types.Error{
			Message: err.Error(),
		}))
		return
	}

	jsonutils.SendJSON(w, common.NoContentResponse())
}
