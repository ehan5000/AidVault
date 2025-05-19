package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type AidRequest struct {
	Name        string `json:"name"`
	AidType     string `json:"aidType"`
	OrgID       int    `json:"orgId"`
	DocumentURL string `json:"documentUrl,omitempty"`
}

var AidDB *sql.DB

func SubmitAidRequest(w http.ResponseWriter, r *http.Request) {
	var req AidRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err := AidDB.Exec(
		"INSERT INTO aid_requests (name, aid_type, org_id, document_url) VALUES ($1, $2, $3, $4)",
		req.Name, req.AidType, req.OrgID, req.DocumentURL,
	)
	if err != nil {
		http.Error(w, "Database insert failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Aid request submitted"))
}

func GetAidStatus(w http.ResponseWriter, r *http.Request) {
	rows, err := AidDB.Query("SELECT id, name, aid_type, status FROM aid_requests")
	if err != nil {
		http.Error(w, "Failed to fetch requests", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []map[string]string
	for rows.Next() {
		var id int
		var name, aidType, status string
		rows.Scan(&id, &name, &aidType, &status)
		results = append(results, map[string]string{
			"id":      fmt.Sprint(id),
			"name":    name,
			"aidType": aidType,
			"status":  status,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func GetRequestStatusByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var name, aidType, status string
	err := AidDB.QueryRow("SELECT name, aid_type, status FROM aid_requests WHERE id = $1", id).
		Scan(&name, &aidType, &status)
	if err != nil {
		http.Error(w, "Request not found", http.StatusNotFound)
		return
	}

	response := map[string]string{
		"name":    name,
		"aidType": aidType,
		"status":  status,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func FulfillAidRequest(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	_, err := AidDB.Exec("UPDATE aid_requests SET status = 'fulfilled' WHERE id = $1", id)
	if err != nil {
		http.Error(w, "Failed to update status", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Request marked as fulfilled"))
}
