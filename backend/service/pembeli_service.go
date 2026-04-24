package service

import (
	"tes-kabayan/backend/models"
	"tes-kabayan/backend/repo"
)

type PembeliService interface {
	GetAllPembeli(search string) ([]models.Pembeli, error)
	GetPembeliByID(id uint) (*models.Pembeli, error)
	CreatePembeli(req models.PembeliReq) (*models.Pembeli, error)
	UpdatePembeli(id uint, req models.PembeliReq) (*models.Pembeli, error)
	DeletePembeli(id uint) error
}

type pembeliService struct {
	repo repo.PembeliRepository
}

func NewPembeliService(repo repo.PembeliRepository) PembeliService {
	return &pembeliService{
		repo: repo,
	}
}

func (s *pembeliService) GetAllPembeli(search string) ([]models.Pembeli, error) {
	return s.repo.FindAll(search)
}

func (s *pembeliService) GetPembeliByID(id uint) (*models.Pembeli, error) {
	return s.repo.FindByID(id)
}

func (s *pembeliService) CreatePembeli(req models.PembeliReq) (*models.Pembeli, error) {
	pembeli := models.Pembeli{
		Nama:   req.Nama,
		Alamat: req.Alamat,
		NoTelp: req.NoTelp,
	}
	err := s.repo.Create(pembeli)
	return &pembeli, err
}

func (s *pembeliService) UpdatePembeli(id uint, req models.PembeliReq) (*models.Pembeli, error) {
	pembeli := models.Pembeli{
		ID:     id,
		Nama:   req.Nama,
		Alamat: req.Alamat,
		NoTelp: req.NoTelp,
	}
	err := s.repo.Update(pembeli)
	return &pembeli, err
}

func (s *pembeliService) DeletePembeli(id uint) error {
	return s.repo.Delete(id)
}
