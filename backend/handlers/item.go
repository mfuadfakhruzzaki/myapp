package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mfuadfakhruzzaki/myapp-backend/models"
	"github.com/mfuadfakhruzzaki/myapp-backend/utils"
)

// GetItemsHandler mengambil semua items dari database.
func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := utils.DB.Query("SELECT id, name, description FROM items")
	if err != nil {
		http.Error(w, "Error fetching items", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Description); err != nil {
			http.Error(w, "Error scanning item", http.StatusInternalServerError)
			return
		}
		items = append(items, item)
	}

	json.NewEncoder(w).Encode(items)
}

// CreateItemHandler menambahkan item baru ke database.
func CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO items (name, description) VALUES ($1, $2) RETURNING id"
	err := utils.DB.QueryRow(query, item.Name, item.Description).Scan(&item.ID)
	if err != nil {
		http.Error(w, "Error creating item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

// UpdateItemHandler mengubah data item yang sudah ada.
func UpdateItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	query := "UPDATE items SET name = $1, description = $2 WHERE id = $3"
	_, err = utils.DB.Exec(query, item.Name, item.Description, id)
	if err != nil {
		http.Error(w, "Error updating item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

// DeleteItemHandler menghapus item dari database.
func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM items WHERE id = $1"
	_, err = utils.DB.Exec(query, id)
	if err != nil {
		http.Error(w, "Error deleting item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
