package routes

import (
    "database/sql"
    "github.com/gorilla/mux"
    "go-core-modules/modules/siswa"
    "go-core-modules/modules/siswa/repositories"
    "go-core-modules/modules/siswa/services"
)

func SetupRoutes(r *mux.Router, db *sql.DB) {
    // Setup siswa module
    siswaRepo := repositories.SiswaRepository{DB: db}
    siswaService := services.SiswaService{Repo: &siswaRepo}
    siswa.SetupRoutes(r, &siswaService)
}