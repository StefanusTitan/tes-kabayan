package service

import (
	"tes-kabayan/backend/models"
	"tes-kabayan/backend/repo"
)

type TransaksiService interface {
	GetAllTransaksi(filter repo.TransaksiFilter) ([]models.Transaksi, error)
	GetTransaksiByID(id uint) (*models.Transaksi, error)
	CreateTransaksi(req models.TransaksiReq) (*models.Transaksi, error)
	DeleteTransaksi(id uint) error
}

type transaksiService struct {
	repo repo.TransaksiRepository
}

func NewTransaksiService(repo repo.TransaksiRepository) TransaksiService {
	return &transaksiService{
		repo: repo,
	}
}

func (s *transaksiService) GetAllTransaksi(filter repo.TransaksiFilter) ([]models.Transaksi, error) {
	// Opsional: bisa tambahkan business rule utk filter tertentu, misal hanya admin yang bisa lihat semua transaksi, pembeli hanya bisa lihat transaksi mereka sendiri
	return s.repo.FindAll(filter)
}

func (s *transaksiService) GetTransaksiByID(id uint) (*models.Transaksi, error) {
	return s.repo.FindByID(id)
}

func (s *transaksiService) CreateTransaksi(req models.TransaksiReq) (*models.Transaksi, error) {
	transaksi := models.Transaksi{
		PembeliID: req.PembeliID,
		BarangID:  req.BarangID,
		Quantity:  req.Quantity,
	}
	err := s.repo.Create(&transaksi)
	return &transaksi, err
}

func (s *transaksiService) DeleteTransaksi(id uint) error {
	return s.repo.Delete(id)
}
