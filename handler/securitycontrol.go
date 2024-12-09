package handler

import (
	"encoding/json"
	"exhttp/models"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (h *AppDB) CreateSecurityControl(w http.ResponseWriter, r *http.Request) {
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

	var seccontrol models.SecurityControl

	// Decode JSON request
	if err := json.NewDecoder(r.Body).Decode(&seccontrol); err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	// Validate input
	if seccontrol.Name == "" {
		http.Error(w, "Name field is required", http.StatusBadRequest)
		return
	}
	// Validate input
	if seccontrol.Risk == "" {
		http.Error(w, "Risk field is required", http.StatusBadRequest)
		return
	}

	// Save to database
	if err := h.DB.Create(&seccontrol).Error; err != nil {
		http.Error(w, "Failed to save service", http.StatusInternalServerError)
		return
	}

	// Respond with the created service
	w.Header().Set("Access-Control-Allow-Origin", "*")                                // Allow all origins (you can restrict this)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Allowed methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,, Authorization")    // Allowed headers (Content-Type)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(seccontrol)
}

func (h *AppDB) ListSecurityControl(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var seccontrols []models.SecurityControl

	// Fetch services with related audits and controls
	err := h.DB.Find(&seccontrols).Error
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
	json.NewEncoder(w).Encode(seccontrols)
}

// RequestPayload defines the structure of the incoming JSON payload
type RequestPayload struct {
	ServiceID  string   `json:"serviceId"`
	ControlIDs []string `json:"controlIds"`
}

// mapControlsHandler handles the POST request
func (h *AppDB) MapControlsHandler(w http.ResponseWriter, r *http.Request) {

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
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse the JSON request body
	var payload RequestPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the ServiceID as a UUID
	serviceUUID, err := uuid.Parse(payload.ServiceID)
	if err != nil {
		http.Error(w, "Invalid ServiceID: must be a valid UUID", http.StatusBadRequest)
		return
	}

	// Validate each ControlID as a UUID
	controlUUIDs := make([]uuid.UUID, len(payload.ControlIDs))
	for i, id := range payload.ControlIDs {
		controlUUID, err := uuid.Parse(id)
		if err != nil {
			http.Error(w, fmt.Sprintf("Invalid ControlID at index %d: must be a valid UUID", i), http.StatusBadRequest)
			return
		}
		controlUUIDs[i] = controlUUID
	}

	// Map controls to the service in the database
	for _, controlUUID := range controlUUIDs {
		serviceControl := models.ServiceSecurityControl{
			ServiceID:         serviceUUID,
			SecurityControlID: controlUUID,
		}

		// Use GORM's `Create` method to save the mapping
		if err := h.DB.Create(&serviceControl).Error; err != nil {
			http.Error(w, fmt.Sprintf("Failed to map control ID %d: %v", controlUUID, err), http.StatusInternalServerError)
			return
		}
	}
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")                                // Allow all origins (you can restrict this)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Allowed methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,, Authorization")    // Allowed headers (Content-Type)
	w.WriteHeader(http.StatusOK)                                                      // Send 200 OK response for OPTIONS

	// Respond with success
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Controls mapped successfully"})
}
