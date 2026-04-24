package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"type:varchar(100);unique;not null" json:"username"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`
}

type UserReq struct {
	Username string `json:"username,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}
