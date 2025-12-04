package products

import (
	"net/http"

	"github.com/whiteblueskyss/golang/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(s Service) *handler {
	return &handler{
		service: s,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	// call the service to get products

	// respond with products in JSON

	//mock product

	products := struct {
		Products []string `json:"products"` // json:"products" tag for proper JSON encoding
	}{}

	json.Write(w, http.StatusOK, products)
}
