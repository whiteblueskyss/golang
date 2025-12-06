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

	products, err := h.service.ListProducts(r.Context())
	if err != nil {
		http.Error(w, "Failed to fetch products", http.StatusInternalServerError)
		return
	}

	productsResponse := struct {
		Products []string `json:"products"`
	}{
		Products: products,
	}

	json.Write(w, http.StatusOK, productsResponse)
}
