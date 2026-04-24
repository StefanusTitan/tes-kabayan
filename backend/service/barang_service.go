package service

import (
	"tes-kabayan/backend/models"
	"tes-kabayan/backend/repo"
)

type BarangService interface {
	GetAllBarang(search string) ([]models.Barang, error)
	GetBarangByID(id uint) (*models.Barang, error)
	CreateBarang(barang *models.Barang) (*models.Barang, error)
	UpdateBarang(id uint, barang *models.Barang) (*models.Barang, error)
	DeleteBarang(id uint) error
}

type barangService struct {
	repo repo.BarangRepository
}

func NewBarangService(repo repo.BarangRepository) BarangService {
	return &barangService{
		repo: repo,
	}
}

func (s *barangService) GetAllBarang(search string) ([]models.Barang, error) {
	return s.repo.FindAll(search)
}

func (s *barangService) GetBarangByID(id uint) (*models.Barang, error) {
	return s.repo.FindByID(id)
}

func (s *barangService) CreateBarang(barang *models.Barang) (*models.Barang, error) {
	err := s.repo.Create(barang)
	return barang, err
}

func (s *barangService) UpdateBarang(id uint, barang *models.Barang) (*models.Barang, error) {
	barang.ID = id
	err := s.repo.Update(barang)
	return barang, err
}

func (s *barangService) DeleteBarang(id uint) error {
	return s.repo.Delete(id)
}
