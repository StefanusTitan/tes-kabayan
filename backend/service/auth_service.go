package service

import (
	"errors"
	"tes-kabayan/backend/models"
	"tes-kabayan/backend/repo"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type AuthService interface {
	Register(req models.UserReq) (*models.User, error)
	Login(req models.UserReq) (string, *models.User, error)
	ValidateToken(tokenStr string) (*Claims, error)
}

type authService struct {
	userRepo  repo.UserRepository
	secretKey string
}

func NewAuthService(userRepo repo.UserRepository, secretKey string) AuthService {
	return &authService{
		userRepo:  userRepo,
		secretKey: secretKey,
	}
}

func (s *authService) Register(req models.UserReq) (*models.User, error) {
	existing, _ := s.userRepo.FindByUsername(req.Username)
	if existing != nil && existing.ID != 0 {
		return nil, errors.New("username sudah digunakan")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: req.Username,
		Password: string(hashed),
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return &models.User{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
	}, nil
}

func (s *authService) Login(req models.UserReq) (string, *models.User, error) {
	user, err := s.userRepo.FindByUsername(req.Username)
	if err != nil {
		return "", nil, errors.New("username atau password salah")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", nil, errors.New("username atau password salah")
	}

	// Generate JWT
	token, err := s.generateToken(user)
	if err != nil {
		return "", nil, err
	}

	return token, &models.User{
		ID:       user.ID,
		Username: user.Username,
		Password: user.Password,
	}, nil
}

func (s *authService) generateToken(user *models.User) (string, error) {
	claims := Claims{
		UserID:   user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "tes-kabayan",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secretKey))
}

func (s *authService) ValidateToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{},
		func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("signing method tidak valid")
			}
			return []byte(s.secretKey), nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("token tidak valid")
	}

	return claims, nil
}
