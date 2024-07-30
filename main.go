package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Athooh/HealthChain/Backend/database"
	handler "github.com/Athooh/HealthChain/handlers"
	"github.com/Athooh/HealthChain/models"
)

func main() {

	if len(os.Args) != 1 {
		fmt.Println("Usage: go run .")
		return
	}

	db, err := database.ConnectDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	// Auto Migrate the schema
	err = db.AutoMigrate(&models.Patient{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	http.HandleFunc("/patient", handler.CreatePatient) // Use POST for creating
	http.HandleFunc("/all", handler.GetAllPatients)    // Use POST for creating
	http.HandleFunc("/patient/", handler.GetPatientHandler)
	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/about", handler.AboutHandler)
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/signup", handler.SignupHandler)
	http.HandleFunc("/signup/facility", handler.SignupFacilityHandler)
	http.HandleFunc("/facility/dashboard", handler.DoctorDashHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	fmt.Println("Server started at http://localhost:8080")

	err1 := http.ListenAndServe(":8080", nil)
	if err1 != nil {
		fmt.Println("Error starting server:", err)
	}
}
