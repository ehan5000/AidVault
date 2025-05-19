package routes

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Organization struct {
	OrgName string `json:"orgName"`
	Email   string `json:"email"`
}

var OrgDB *sql.DB

func RegisterOrganization(w http.ResponseWriter, r *http.Request) {
	var org Organization
	if err := json.NewDecoder(r.Body).Decode(&org); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err := OrgDB.Exec("INSERT INTO organizations (org_name, email) VALUES ($1, $2)", org.OrgName, org.Email)
	if err != nil {
		http.Error(w, "Database insert failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Organization registered"))
}

func GetAllOrganizations(w http.ResponseWriter, r *http.Request) {
	rows, err := OrgDB.Query("SELECT id, org_name, email FROM organizations")
	if err != nil {
		http.Error(w, "Failed to fetch organizations", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []map[string]interface{}
	for rows.Next() {
		var id int
		var name, email string
		rows.Scan(&id, &name, &email)
		results = append(results, map[string]interface{}{
			"id":    id,
			"name":  name,
			"email": email,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
