package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var categories = []Category{
	{ID: 1, Name: "Elektronik", Description: "Gadget dan alat listrik"},
	{ID: 2, Name: "Pakaian", Description: "Baju, celana, dan aksesoris"},
}

func handleCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(categories)
	} else if r.Method == "POST" {
		var newCategory Category
		if err := json.NewDecoder(r.Body).Decode(&newCategory); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		newCategory.ID = len(categories) + 1
		categories = append(categories, newCategory)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newCategory)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleCategoryByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	index := -1
	for i, c := range categories {
		if c.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(categories[index])

	case "PUT":
		var updateData Category
		if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		updateData.ID = id
		categories[index] = updateData
		json.NewEncoder(w).Encode(updateData)

	case "DELETE":
		categories = append(categories[:index], categories[index+1:]...)
		json.NewEncoder(w).Encode(map[string]string{"message": "Category deleted successfully"})

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/categories", handleCategories)
	http.HandleFunc("/categories/", handleCategoryByID)

	fmt.Println("Server running at localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
