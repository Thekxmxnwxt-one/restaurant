package service

import (
	"context"
	"errors"
	"os"
	"time"

	"restaurant/internal/models"
	"restaurant/internal/repository"
	"restaurant/internal/utils"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	db        repository.AuthRepository
	jwtSecret []byte
}

func NewAuthApp(db repository.AuthRepository) *AuthService {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret" // ควรตั้งค่าจริงใน env เท่านั้น
	}
	return &AuthService{db: db, jwtSecret: []byte(secret)}
}

func (app *AuthService) Authenticate(ctx context.Context, username, password string) (string, error) {
	user, err := app.db.GetUserByUsername(ctx, username)
	if err != nil {
		return "", errors.New("user not found")
	}

	// ตรวจสอบรหัสผ่าน
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("รหัสผ่านไม่ถูกต้อง")
	}

	// สร้าง JWT token
	token, err := app.generateJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (app *AuthService) generateJWT(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(app.jwtSecret)
}

func (app *AuthService) RegisterUser(ctx context.Context, name, username, password, role string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	var userID int
	if role == "staff" {
		userID, err = app.db.CreateEmployee(ctx, name, username, string(hashed))
	} else if role == "customer" {
		userID, err = app.db.CreateCustomer(ctx, name, username, string(hashed))
	} else {
		return "", errors.New("role ต้องเป็น 'staff' หรือ 'customer'")
	}

	if err != nil {
		return "", err
	}

	return utils.GenerateJWT(userID, username, role)
}
