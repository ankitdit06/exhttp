package handler

import (
	"encoding/json"
	"exhttp/models"
	"net/http"
	"strings"
)

func (h *AppDB) FetchControlsByServiceId(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}
	// Retrieve the "username" query parameter
	path := r.URL.Path

	// Check if the path starts with "/get/"
	if !strings.HasPrefix(path, "/Service/SecurityControl/") {
		http.Error(w, "Invalid URL path", http.StatusNotFound)
		return
	}
	// Extract the username by trimming the prefix
	sid := strings.TrimPrefix(path, "/Service/SecurityControl/")

	var ssc []models.ServiceSecurityControl

	// Fetch services with related audits and controls
	err := h.DB.Preload("SecurityControl").Where("service_id = ?", sid).Find(&ssc).Error
	if err != nil {
		http.Error(w, "Failed to fetch services", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")                                // Allow all origins
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Allowed methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")     // Allowed headers

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ssc)
}
