package handler

import (
	"net/http"
	"strconv"
	"tes-kabayan/backend/models"
	"tes-kabayan/backend/repo"
	"tes-kabayan/backend/service"
	"time"

	"github.com/gin-gonic/gin"
)

type TransaksiHandler struct {
	svc service.TransaksiService
}

func NewTransaksiHandler(svc service.TransaksiService) *TransaksiHandler {
	return &TransaksiHandler{svc: svc}
}

func (h *TransaksiHandler) GetAllTransaksi(c *gin.Context) {
	var filter repo.TransaksiFilter
	filter.Search = c.Query("search") // ?search=baju

	if pembeliID := c.Query("pembeli_id"); pembeliID != "" {
		id, err := strconv.ParseUint(pembeliID, 10, 64)
		if err == nil {
			filter.PembeliID = uint(id)
		}
	}

	// expected format from frontend: 2024-01-01
	if start := c.Query("start_date"); start != "" {
		t, err := time.Parse("2006-01-02", start)
		if err == nil {
			filter.StartDate = t
		}
	}

	if end := c.Query("end_date"); end != "" {
		t, err := time.Parse("2006-01-02", end)
		if err == nil {
			filter.EndDate = t
		}
	}

	transaksi, err := h.svc.GetAllTransaksi(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, transaksi)
}

func (h *TransaksiHandler) GetTransaksiByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	transaksi, err := h.svc.GetTransaksiByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaksi not found"})
		return
	}
	c.JSON(http.StatusOK, transaksi)
}

func (h *TransaksiHandler) CreateTransaksi(c *gin.Context) {
	var req models.TransaksiReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	transaksi, err := h.svc.CreateTransaksi(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, transaksi)
}

func (h *TransaksiHandler) DeleteTransaksi(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	err = h.svc.DeleteTransaksi(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaksi deleted"})
}
