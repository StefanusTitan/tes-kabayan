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
	config.DB.AutoMigrate(&models.User{}, &models.Barang{}, &models.Pembeli{}, &models.Transaksi{})

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
