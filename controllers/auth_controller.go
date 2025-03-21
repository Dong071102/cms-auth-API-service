package controllers

import (
	"go-cms-backend/config"
	"go-cms-backend/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	pgvector "github.com/pgvector/pgvector-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtSecret = []byte("your-secret-key")

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func GenerateFakeEmbedding(dim int) pgvector.Vector {
	vec := pgvector.NewVector(make([]float32, dim))
	return vec
}
func generateJWT(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.UserID.String(),
		"role":    user.Role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func RegisterUser(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	var existing models.User
	if err := config.DB.Where("username = ? OR email = ?", user.Username, user.Email).First(&existing).Error; err == nil {
		return c.JSON(http.StatusConflict, map[string]string{"message": "Username or Email already exists"})
	}

	hashed, err := HashPassword(user.PasswordHash)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Password hashing failed"})
	}
	user.PasswordHash = hashed

	config.DB.Create(&user)

	switch user.Role {
	case "student":
		var studentCode string
		for {
			code := models.GenerateStudentCode()
			var count int64
			config.DB.Model(&models.Student{}).Where("student_code = ?", code).Count(&count)
			if count == 0 {
				studentCode = code
				break
			}
		}
		s := models.Student{
			StudentID:     user.UserID,
			StudentCode:   studentCode,
			FaceEmbedding: GenerateFakeEmbedding(512)}
		config.DB.Create(&s)
	case "lecturer":
		l := models.Lecturer{LecturerID: user.UserID, LectainerCode: uuid.New().String()}
		config.DB.Create(&l)
	case "admin":
		a := models.Admin{AdminID: user.UserID, AdminCode: uuid.New().String()}
		config.DB.Create(&a)
	default:
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid role"})
	}

	return c.JSON(http.StatusCreated, map[string]string{
		"message": "User registered successfully",
		"user_id": user.UserID.String(),
	})
}

func LoginUser(c echo.Context) error {
	var input struct {
		UsernameOrEmail string `json:"username_or_email"`
		Password        string `json:"password"`
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	var user models.User
	if err := config.DB.Where("username = ? OR email = ?", input.UsernameOrEmail, input.UsernameOrEmail).
		First(&user).Error; err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "User not found"})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Incorrect password"})
	}

	token, err := generateJWT(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Login successful",
		"token":   token,
	})
}
