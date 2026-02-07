package handlers

import (
	"encoding/json"
	"go-crud/models"
	"go-crud/services"
	"net/http"
	"strconv"
	"strings"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) HandleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		name := r.URL.Query().Get("name")

		products, err := h.service.GetAll(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)

	case http.MethodPost:
		var product models.Product
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			http.Error(w, "Invalid Payload", http.StatusBadRequest)
			return
		}
		if err := h.service.Create(&product); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(product)
	}
}

func (h *ProductHandler) HandleProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/products/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodPut:
		var product models.Product
		json.NewDecoder(r.Body).Decode(&product)
		h.service.Update(id, &product)
		json.NewEncoder(w).Encode(product)

	case http.MethodDelete:
		h.service.Delete(id)
		json.NewEncoder(w).Encode(map[string]string{"message": "deleted"})
	}
}
