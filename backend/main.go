package main

import (
	"os"
	"tes-kabayan/backend/config"
	"tes-kabayan/backend/handler"
	"tes-kabayan/backend/models"
	"tes-kabayan/backend/repo"
	"tes-kabayan/backend/routes"
	"tes-kabayan/backend/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm/clause"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	db_name := os.Getenv("DB_NAME")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	config.InitDB(db_name, db_user, db_password, db_host, db_port)
	config.DB.AutoMigrate(&models.User{}, &models.Barang{}, &models.Pembeli{})

	// Auth
	userRepo := repo.NewUserRepository(config.DB)
	authService := service.NewAuthService(userRepo, os.Getenv("SECRET_KEY"))
	authHandler := handler.NewAuthHandler(authService)

	// Barang
	barangRepo := repo.NewBarangRepository(config.DB)
	barangService := service.NewBarangService(barangRepo)
	barangHandler := handler.NewBarangHandler(barangService)

	// Pembeli
	pembeliRepo := repo.NewPembeliRepository(config.DB)
	pembeliService := service.NewPembeliService(pembeliRepo)
	pembeliHandler := handler.NewPembeliHandler(pembeliService)

	// Transaksi
	transaksiRepo := repo.NewTransaksiRepository(config.DB)
	transaksiService := service.NewTransaksiService(transaksiRepo)
	transaksiHandler := handler.NewTransaksiHandler(transaksiService)

	r := gin.Default()

	routes.SetupRoutes(authHandler, pembeliHandler, barangHandler, transaksiHandler, authService, r)

	r.Run(":" + os.Getenv("PORT"))
}

func seedData() {
	user := models.User{
		Username: "admin",
		Password: "admin123",
	}

	barang := []models.Barang{
		{
			Nama:      "Barang 1",
			Deskripsi: "Barang 1 yang sangat bagus",
			Harga:     10000,
			Stock:     10,
		},
		{
			Nama:      "Barang 2",
			Deskripsi: "Barang 2 yang sangat bagus",
			Harga:     20000,
			Stock:     20,
		},
		{
			Nama:      "Barang 3",
			Deskripsi: "Barang 3 yang sangat bagus",
			Harga:     30000,
			Stock:     30,
		},
	}

	pembeli := []models.Pembeli{
		{
			Nama:   "Pembeli 1",
			Alamat: "Jl. Pembeli 1 No. 1",
			NoTelp: "081234567890",
		},
		{
			Nama:   "Pembeli 2",
			Alamat: "Jl. Pembeli 2 No. 2",
			NoTelp: "081234567891",
		},
		{
			Nama:   "Pembeli 3",
			Alamat: "Jl. Pembeli 3 No. 3",
			NoTelp: "081234567892",
		},
	}

	config.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&user)

	config.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&barang)

	config.DB.Clauses(clause.OnConflict{DoNothing: true}).Create(&pembeli)
}
