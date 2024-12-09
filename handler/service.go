package handler

import (
	"encoding/json"
	"exhttp/models"
	"fmt"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

// Local wrapper for the imported APPDB type
type AppDB struct {
	DB *gorm.DB
}

func (h *AppDB) CreateService(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")                                // Allow all origins (you can restrict this)
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Allowed methods
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,, Authorization")    // Allowed headers (Content-Type)
		w.WriteHeader(http.StatusOK)                                                      // Send 200 OK response for OPTIONS
		fmt.Fprintf(w, "OPTIONS request handled")
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var service models.Service

	// Decode JSON request
	if err := json.NewDecoder(r.Body).Decode(&service); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	// Validate input
	if service.Name == "" {
		http.Error(w, "Name field is required", http.StatusBadRequest)
		return
	}

	// Save to database
	if err := h.DB.Create(&service).Error; err != nil {
		http.Error(w, "Failed to save service", http.StatusInternalServerError)
		return
	}

	// Respond with the created service
	w.Header().Set("Access-Control-Allow-Origin", "*")                                // Allow all origins (you can restrict this)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Allowed methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,, Authorization")    // Allowed headers (Content-Type)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(service)
}

func (h *AppDB) ListServices(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var services []models.Service

	// Fetch services with related audits and controls
	err := h.DB.Preload("Team").Preload("SecAudit").Find(&services).Error
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
	json.NewEncoder(w).Encode(services)
}

func (h *AppDB) FetchService(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}
	// Retrieve the "username" query parameter
	path := r.URL.Path

	// Check if the path starts with "/get/"
	if !strings.HasPrefix(path, "/Service/") {
		http.Error(w, "Invalid URL path", http.StatusNotFound)
		return
	}
	// Extract the username by trimming the prefix
	sid := strings.TrimPrefix(path, "/Service/")

	var service models.Service

	// Fetch services with related audits and controls
	err := h.DB.Preload("Team").Preload("SecAudit").First(&service, "id=?", sid).Error
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
	json.NewEncoder(w).Encode(service)
}
