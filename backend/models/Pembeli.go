package models

type Pembeli struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Nama   string `gorm:"type:varchar(100), not null" json:"nama"`
	Alamat string `gorm:"type:varchar(255)" json:"alamat"`
	NoTelp string `gorm:"type:varchar(20)" json:"no_telp"`
}
