package services

import (
    "go-core-modules/modules/siswa/models"
    "go-core-modules/modules/siswa/repositories"
)

type SiswaService struct {
    Repo *repositories.SiswaRepository
}

func (s *SiswaService) GetSiswaByID(id int) (*models.Siswa, error) {
    return s.Repo.GetByID(id)
}