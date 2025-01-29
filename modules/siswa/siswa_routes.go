package siswa

import (
    "github.com/gorilla/mux"
    "go-core-modules/modules/siswa/handlers"
    "go-core-modules/modules/siswa/services"
)

func SetupRoutes(r *mux.Router, service *services.SiswaService) {
    handler := &handlers.SiswaHandler{Service: service}
    r.HandleFunc("/siswa/{id}", handler.GetSiswaByID).Methods("GET")
}