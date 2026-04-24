package models

type Pembelian struct {
	ID        uint `gorm:"primaryKey" json:"id"`
	PembeliID uint `gorm:"not null" json:"pembeli_id"`
	BarangID  uint `gorm:"not null" json:"barang_id"`
	Quantity  int  `gorm:"not null" json:"quantity"`
}
