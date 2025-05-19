package main

import (
	"aidvault/db"
	"aidvault/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	err = db.InitTables(database)
	if err != nil {
		log.Fatal("Table setup failed:", err)
	}

	routes.OrgDB = database
	routes.AidDB = database

	r := mux.NewRouter()
	r.HandleFunc("/aid-request", routes.SubmitAidRequest).Methods("POST")
	r.HandleFunc("/aid-request/{id}/status", routes.GetRequestStatusByID).Methods("GET")
	r.HandleFunc("/aid-request/{id}/fulfill", routes.FulfillAidRequest).Methods("PATCH")
	r.HandleFunc("/aid-status", routes.GetAidStatus).Methods("GET")
	r.HandleFunc("/org", routes.RegisterOrganization).Methods("POST")
	r.HandleFunc("/org/all", routes.GetAllOrganizations).Methods("GET")

	fmt.Println("AidVault API running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
