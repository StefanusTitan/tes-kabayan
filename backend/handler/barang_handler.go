package handler

import (
	"net/http"
	"strconv"
	"tes-kabayan/backend/models"
	"tes-kabayan/backend/service"

	"github.com/gin-gonic/gin"
)

type BarangHandler struct {
	svc service.BarangService
}

func NewBarangHandler(svc service.BarangService) *BarangHandler {
	return &BarangHandler{svc: svc}
}

func (h *BarangHandler) GetAllBarang(c *gin.Context) {
	search := c.Query("search")
	barang, err := h.svc.GetAllBarang(search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, barang)
}

func (h *BarangHandler) GetBarangByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	barang, err := h.svc.GetBarangByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Barang not found"})
		return
	}
	c.JSON(http.StatusOK, barang)
}

func (h *BarangHandler) CreateBarang(c *gin.Context) {
	var req models.Barang
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	barang, err := h.svc.CreateBarang(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, barang)
}

func (h *BarangHandler) UpdateBarang(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var req models.Barang
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	barang, err := h.svc.UpdateBarang(uint(id), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, barang)
}

func (h *BarangHandler) DeleteBarang(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	err = h.svc.DeleteBarang(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Barang deleted"})
}
