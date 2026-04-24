package repo

import (
	"tes-kabayan/backend/models"

	"gorm.io/gorm"
)

type PembeliRepository interface {
	FindAll(search string) ([]models.Pembeli, error)
	FindByID(id uint) (*models.Pembeli, error)
	Create(pembeli models.Pembeli) error
	Update(pembeli models.Pembeli) error
	Delete(id uint) error
}

type pembeliRepo struct {
	db *gorm.DB
}

func NewPembeliRepository(db *gorm.DB) PembeliRepository {
	return &pembeliRepo{db}
}

func (r *pembeliRepo) FindAll(search string) ([]models.Pembeli, error) {
	var pembeli []models.Pembeli
	query := r.db.Preload("Barang")

	if search != "" {
		query = query.Where("nama LIKE ?", "%"+search+"%")
	}

	err := query.Find(&pembeli).Error
	return pembeli, err
}

func (r *pembeliRepo) FindByID(id uint) (*models.Pembeli, error) {
	var pembeli models.Pembeli
	err := r.db.Preload("Barang").First(&pembeli, id).Error
	return &pembeli, err
}

func (r *pembeliRepo) Create(pembeli models.Pembeli) error {
	return r.db.Create(&pembeli).Error
}

func (r *pembeliRepo) Update(pembeli models.Pembeli) error {
	return r.db.Model(&models.Pembeli{}).Where("id = ?", pembeli.ID).Updates(pembeli).Error
}

func (r *pembeliRepo) Delete(id uint) error {
	return r.db.Delete(&models.Pembeli{}, id).Error
}
