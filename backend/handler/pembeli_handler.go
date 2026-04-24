package handler

import (
	"net/http"
	"strconv"
	"tes-kabayan/backend/models"
	"tes-kabayan/backend/service"

	"github.com/gin-gonic/gin"
)

type PembeliHandler struct {
	svc service.PembeliService
}

func NewPembeliHandler(svc service.PembeliService) *PembeliHandler {
	return &PembeliHandler{svc: svc}
}

func (h *PembeliHandler) GetAllPembeli(c *gin.Context) {
	search := c.Query("search")
	pembeli, err := h.svc.GetAllPembeli(search)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pembeli)
}

func (h *PembeliHandler) GetPembeliByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	pembeli, err := h.svc.GetPembeliByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pembeli not found"})
		return
	}
	c.JSON(http.StatusOK, pembeli)
}

func (h *PembeliHandler) CreatePembeli(c *gin.Context) {
	var req models.PembeliReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pembeli, err := h.svc.CreatePembeli(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, pembeli)
}

func (h *PembeliHandler) UpdatePembeli(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var req models.PembeliReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pembeli, err := h.svc.UpdatePembeli(uint(id), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pembeli)
}

func (h *PembeliHandler) DeletePembeli(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	err = h.svc.DeletePembeli(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Pembeli deleted"})
}
