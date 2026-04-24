package repo

import (
	"tes-kabayan/backend/models"

	"gorm.io/gorm"
)

type BarangRepository interface {
	FindAll(search string) ([]models.Barang, error)
	FindByID(id uint) (*models.Barang, error)
	Create(barang models.Barang) error
	Update(barang models.Barang) error
	Delete(id uint) error
}

type barangRepo struct {
	db *gorm.DB
}

func NewBarangRepository(db *gorm.DB) BarangRepository {
	return &barangRepo{db}
}

func (r *barangRepo) FindAll(search string) ([]models.Barang, error) {
	var barang []models.Barang
	query := r.db.Model(&models.Barang{})
	if search != "" {
		query = query.Where("nama LIKE ?", "%"+search+"%")
	}
	err := query.Find(&barang).Error
	return barang, err
}

func (r *barangRepo) FindByID(id uint) (*models.Barang, error) {
	var barang models.Barang
	err := r.db.First(&barang, id).Error
	return &barang, err
}

func (r *barangRepo) Create(barang models.Barang) error {
	return r.db.Create(&barang).Error
}

func (r *barangRepo) Update(barang models.Barang) error {
	return r.db.Model(&models.Barang{}).Where("id = ?", barang.ID).Updates(barang).Error
}

func (r *barangRepo) Delete(id uint) error {
	return r.db.Delete(&models.Barang{}, id).Error
}
