package repositories

import (
    "database/sql"
    "go-core-modules/modules/siswa/models"
)

type SiswaRepository struct {
    DB *sql.DB
}

func (r *SiswaRepository) GetByID(id int) (*models.Siswa, error) {
    var siswa models.Siswa
    err := r.DB.QueryRow("SELECT id, nama, kelas FROM siswa WHERE id = ?", id).Scan(&siswa.ID, &siswa.Nama, &siswa.Kelas)
    if err != nil {
        return nil, err
    }
    return &siswa, nil
}