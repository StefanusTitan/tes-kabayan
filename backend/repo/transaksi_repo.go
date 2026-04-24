package repo

import (
	"strings"
	"tes-kabayan/backend/models"
	"time"

	"gorm.io/gorm"
)

type TransaksiRepository interface {
	FindAll(filter TransaksiFilter) ([]models.Transaksi, error)
	FindByID(id uint) (*models.Transaksi, error)
	Create(transaksi *models.Transaksi) error
	Delete(id uint) error
}

type transaksiRepo struct {
	db *gorm.DB
}

func NewTransaksiRepository(db *gorm.DB) TransaksiRepository {
	return &transaksiRepo{db}
}

type TransaksiFilter struct {
	Search    string // Pencarian berdasarkan nama barang atau pembeli
	PembeliID uint   // Kasus admin ingin melihat transaksi berdasarkan pembeli tertentu
	StartDate time.Time
	EndDate   time.Time
}

func (r *transaksiRepo) FindAll(filter TransaksiFilter) ([]models.Transaksi, error) {
	var transaksi []models.Transaksi
	query := r.db.Model(&models.Transaksi{}).
		Preload("Pembeli").
		Preload("Barang")

	if filter.PembeliID != 0 {
		query = query.Where("pembeli_id = ?", filter.PembeliID)
	}

	if !filter.StartDate.IsZero() {
		query = query.Where("created_at >= ?", filter.StartDate)
	}

	if !filter.EndDate.IsZero() {
		query = query.Where("created_at < ?", filter.EndDate.AddDate(0, 0, 1))
	}

	err := query.Order("created_at DESC").Find(&transaksi).Error
	if err != nil {
		return transaksi, err
	}

	if filter.Search == "" {
		return transaksi, nil
	}

	needle := strings.ToLower(strings.TrimSpace(filter.Search))
	filtered := make([]models.Transaksi, 0, len(transaksi))
	for _, t := range transaksi {
		if strings.Contains(strings.ToLower(t.Barang.Nama), needle) || strings.Contains(strings.ToLower(t.Pembeli.Nama), needle) {
			filtered = append(filtered, t)
		}
	}

	return filtered, nil
}

func (r *transaksiRepo) FindByID(id uint) (*models.Transaksi, error) {
	var transaksi models.Transaksi
	err := r.db.First(&transaksi, id).Error
	return &transaksi, err
}

func (r *transaksiRepo) Create(transaksi *models.Transaksi) error {
	return r.db.Create(transaksi).Error
}

func (r *transaksiRepo) Delete(id uint) error {
	return r.db.Delete(&models.Transaksi{}, id).Error
}
