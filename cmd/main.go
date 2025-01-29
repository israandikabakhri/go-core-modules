package main

import (
    "database/sql"
    "log"
    "net/http"
    "go-core-modules/config"
    "go-core-modules/routes"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
)

func main() {
    // Load configuration from .env
    cfg := config.LoadConfig()

    // Connect to database
    db, err := sql.Open("mysql", cfg.DatabaseURL)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer db.Close()

    // Initialize router
    r := mux.NewRouter()

    // Setup routes
    routes.SetupRoutes(r, db)

    // Start server
    log.Println("Server started on :8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}