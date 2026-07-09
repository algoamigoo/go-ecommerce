package products

import (
	"net/http"
	"github.com/algoamigoo/go-ecommerce/internal/json"
)

type handler struct {
	service Service
}

func NewHandler (service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request){
	// 1. call the service ->ListProduct
	// 2. Return Json in an HTTP response
	products := struct{
	 Products []string `json:"products"`
	}{}
	
	json.Write(w, http.StatusOK, products)
}