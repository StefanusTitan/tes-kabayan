package models

import "time"

type Transaksi struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	PembeliID uint      `gorm:"not null;index" json:"pembeli_id"`
	BarangID  uint      `gorm:"not null;index" json:"barang_id"`
	Quantity  int       `gorm:"not null" json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	Pembeli   Pembeli   `gorm:"foreignKey:PembeliID" json:"pembeli"`
	Barang    Barang    `gorm:"foreignKey:BarangID" json:"barang"`
}

type TransaksiReq struct {
	PembeliID uint `json:"pembeli_id,omitempty" binding:"required"`
	BarangID  uint `json:"barang_id,omitempty" binding:"required"`
	Quantity  int  `json:"quantity,omitempty" binding:"required,min=1"`
}
