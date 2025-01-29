package handlers

import (
    "net/http"
    "strconv"
    "go-core-modules/core/utils"
    "go-core-modules/modules/siswa/services"
    "github.com/gorilla/mux"
)

type SiswaHandler struct {
    Service *services.SiswaService
}

func (h *SiswaHandler) GetSiswaByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        utils.JSONError(w, http.StatusBadRequest, "Invalid ID")
        return
    }

    siswa, err := h.Service.GetSiswaByID(id)
    if err != nil {
        utils.JSONError(w, http.StatusInternalServerError, err.Error())
        return
    }

    utils.JSONResponse(w, http.StatusOK, siswa)
}