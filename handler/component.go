package handler

import (
	"encoding/json"
	"exhttp/models"
	"fmt"
	"net/http"
	"strings"
)

func (h *AppDB) CreateComponent(w http.ResponseWriter, r *http.Request) {
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

	var component models.Component

	// Decode JSON request
	if err := json.NewDecoder(r.Body).Decode(&component); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	// Validate input
	if component.Name == "" {
		http.Error(w, "Name field is required", http.StatusBadRequest)
		return
	}

	// Save to database
	if err := h.DB.Create(&component).Error; err != nil {
		http.Error(w, "Failed to save service", http.StatusInternalServerError)
		return
	}

	// Respond with the created service
	w.Header().Set("Access-Control-Allow-Origin", "*")                                // Allow all origins (you can restrict this)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Allowed methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,, Authorization")    // Allowed headers (Content-Type)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(component)
}

func (h *AppDB) ListComponents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var components []models.Component

	// Fetch services with related audits and controls
	err := h.DB.Preload("Team").Find(&components).Error
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
	json.NewEncoder(w).Encode(components)
}

func (h *AppDB) FetchComponentByService(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}
	// Retrieve the "username" query parameter
	path := r.URL.Path

	// Check if the path starts with "/get/"
	if !strings.HasPrefix(path, "/Component/Service/") {
		http.Error(w, "Invalid URL path", http.StatusNotFound)
		return
	}
	// Extract the username by trimming the prefix
	sid := strings.TrimPrefix(path, "/Component/Service/")

	var components []models.Component

	// Fetch services with related audits and controls
	err := h.DB.Preload("Team").Where("service_id = ?", sid).Find(&components).Error
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
	json.NewEncoder(w).Encode(components)
}


func (h *AppDB) FetchComponentById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}
	// Retrieve the "username" query parameter
	path := r.URL.Path

	// Check if the path starts with "/get/"
	if !strings.HasPrefix(path, "/Component/") {
		http.Error(w, "Invalid URL path", http.StatusNotFound)
		return
	}
	// Extract the username by trimming the prefix
	sid := strings.TrimPrefix(path, "/Component/")

	var component models.Component

	// Fetch services with related audits and controls
	err := h.DB.Preload("Team").Where("id = ?", sid).Find(&component).Error
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
	json.NewEncoder(w).Encode(component)
}
