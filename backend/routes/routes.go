package routes

import (
	"tes-kabayan/backend/handler"
	"tes-kabayan/backend/middleware"
	"tes-kabayan/backend/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	authHandler *handler.AuthHandler,
	pembeliHandler *handler.PembeliHandler,
	barangHandler *handler.BarangHandler,
	transaksiHandler *handler.TransaksiHandler,
	authService service.AuthService,
	r *gin.Engine,
) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	api := r.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
		auth.POST("/logout", authHandler.Logout)
	}

	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware(authService))
	{
		protected.GET("/barang", barangHandler.GetAllBarang)
		protected.GET("/barang/:id", barangHandler.GetBarangByID)
		protected.POST("/barang", barangHandler.CreateBarang)
		protected.PUT("/barang/:id", barangHandler.UpdateBarang)
		protected.DELETE("/barang/:id", barangHandler.DeleteBarang)
		protected.GET("/pembeli", pembeliHandler.GetAllPembeli)
		protected.GET("/pembeli/:id", pembeliHandler.GetPembeliByID)
		protected.POST("/pembeli", pembeliHandler.CreatePembeli)
		protected.PUT("/pembeli/:id", pembeliHandler.UpdatePembeli)
		protected.DELETE("/pembeli/:id", pembeliHandler.DeletePembeli)
		protected.GET("/transaksi", transaksiHandler.GetAllTransaksi)
		protected.GET("/transaksi/:id", transaksiHandler.GetTransaksiByID)
		protected.POST("/transaksi", transaksiHandler.CreateTransaksi)
		protected.DELETE("/transaksi/:id", transaksiHandler.DeleteTransaksi)
	}
}
