package repo

import (
	"tes-kabayan/backend/models"
	"time"

	"gorm.io/gorm"
)

type transaksiRepository interface {
	FindAll(filter TransaksiFilter) ([]models.Transaksi, error)
	FindByID(id uint) (*models.Transaksi, error)
	Create(transaksi models.Transaksi) error
	Delete(id uint) error
}

type transaksiRepo struct {
	db *gorm.DB
}

func NewTransaksiRepository(db *gorm.DB) transaksiRepository {
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

	if filter.Search != "" {
		query = query.
			Joins("JOIN barang ON barang.id = transaksi.barang_id").
			Joins("JOIN pembeli ON pembeli.id = transaksi.pembeli_id").
			Where("barang.nama LIKE ? OR pembeli.nama LIKE ?",
				"%"+filter.Search+"%", "%"+filter.Search+"%")
	}

	if filter.PembeliID != 0 {
		query = query.Where("transaksi.pembeli_id = ?", filter.PembeliID)
	}

	if !filter.StartDate.IsZero() {
		query = query.Where("transaksi.created_at >= ?", filter.StartDate)
	}

	if !filter.EndDate.IsZero() {
		query = query.Where("transaksi.created_at < ?", filter.EndDate.AddDate(0, 0, 1))
	}

	err := query.Find(&transaksi).Error
	return transaksi, err
}

func (r *transaksiRepo) FindByID(id uint) (*models.Transaksi, error) {
	var transaksi models.Transaksi
	err := r.db.First(&transaksi, id).Error
	return &transaksi, err
}

func (r *transaksiRepo) Create(transaksi models.Transaksi) error {
	return r.db.Create(&transaksi).Error
}

func (r *transaksiRepo) Delete(id uint) error {
	return r.db.Delete(&models.Transaksi{}, id).Error
}
