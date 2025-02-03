package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mfuadfakhruzzaki/myapp-backend/models"
	"github.com/mfuadfakhruzzaki/myapp-backend/utils"
)

// GetItemsHandler mengambil semua items dari database
func GetItemsHandler(w http.ResponseWriter, r *http.Request) {
	var items []models.Item
	if result := utils.DB.Find(&items); result.Error != nil {
		http.Error(w, "Error fetching items: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(items)
}

// CreateItemHandler membuat item baru dan menyimpannya di database
func CreateItemHandler(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if result := utils.DB.Create(&item); result.Error != nil {
		http.Error(w, "Error creating item: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

// UpdateItemHandler memperbarui data item yang sudah ada berdasarkan ID
func UpdateItemHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	var item models.Item
	// Cari item berdasarkan ID
	if result := utils.DB.First(&item, id); result.Error != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	// Decode data baru dari request body
	var updatedData models.Item
	if err := json.NewDecoder(r.Body).Decode(&updatedData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Perbarui field yang diizinkan
	item.Name = updatedData.Name
	item.Description = updatedData.Description

	// Simpan perubahan
	if result := utils.DB.Save(&item); result.Error != nil {
		http.Error(w, "Error updating item: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(item)
}

// DeleteItemHandler menghapus item berdasarkan ID
func DeleteItemHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	if result := utils.DB.Delete(&models.Item{}, id); result.Error != nil {
		http.Error(w, "Error deleting item: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
