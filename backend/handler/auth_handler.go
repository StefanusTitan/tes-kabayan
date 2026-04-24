package handler

import (
	"net/http"
	"tes-kabayan/backend/models"
	"tes-kabayan/backend/service"
	"time"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct{ svc service.AuthService }

func NewAuthHandler(svc service.AuthService) *AuthHandler {
	return &AuthHandler{svc}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req models.UserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := h.svc.Register(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req models.UserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, user, err := h.svc.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "email atau password salah"})
		return
	}

	c.SetCookie(
		"auth_token",                // nama cookie
		token,                       // value
		int(24*time.Hour.Seconds()), // max age (detik)
		"/",                         // path
		"",                          // domain (kosong = current domain)
		false,                       // secure (true di production/HTTPS)
		true,                        // httpOnly
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login berhasil",
		"data":    user,
	})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	c.SetCookie(
		"auth_token",
		"", // kosongkan value
		-1, // max age -1 = hapus cookie
		"/",
		"",
		false,
		true,
	)
	c.JSON(http.StatusOK, gin.H{"message": "Logout berhasil"})
}
