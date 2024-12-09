package handler

import (
	"encoding/json"
	"exhttp/models"
	"net/http"
)

func (h *AppDB) DashboardCount(w http.ResponseWriter, r *http.Request) {
  var serviceCount, componentCount, teamCount, securityControlCount int64

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}


  // Count records in each table
		h.DB.Model(&models.Service{}).Count(&serviceCount)
		h.DB.Model(&models.Component{}).Count(&componentCount)
		h.DB.Model(&models.Team{}).Count(&teamCount)
		h.DB.Model(&models.SecurityControl{}).Count(&securityControlCount)


    // Prepare response
  		counts := map[string]int64{
  			"serviceCount":        serviceCount,
  			"componentCount":      componentCount,
  			"teamCount":           teamCount,
  			"securityControlCount": securityControlCount,
  		}

	w.Header().Set("Access-Control-Allow-Origin", "*")                                // Allow all origins
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS") // Allowed methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")     // Allowed headers

	// Send JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(counts)
}
