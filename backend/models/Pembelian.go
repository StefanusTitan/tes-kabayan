package models

type Barang struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Nama      string `gorm:"type:varchar(100), not null" json:"nama"`
	Deskripsi string `gorm:"type:text" json:"deskripsi"`
	Harga     int64  `gorm:"not null" json:"harga"`
	Stock     int    `gorm:"not null" json:"stock"`
}
