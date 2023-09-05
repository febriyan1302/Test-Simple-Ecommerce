package middleware

import (
	"fmt"
	"strings"
	"time"

	"github.com/febriyan1302/catalog-service/internal/entity"
	"github.com/febriyan1302/catalog-service/internal/model"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	jtoken "github.com/golang-jwt/jwt/v4"
)

func NewAuthMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(secret),
	})
}

func CreateToken(user model.User) (string, error) {
	day := time.Hour * 24

	claims := jtoken.MapClaims{
		"ID":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(day * 1).Unix(),
	}

	// Create token
	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)

	// Generate Encoded Token
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return t, nil
}

func DeserializeUser(c *fiber.Ctx) error {
	var tokenString string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("token") != "" {
		tokenString = c.Cookies("token")
	}

	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(entity.Response{
			Code:    c.Response().StatusCode(),
			Message: "You are not Authenticate",
			Data:    nil,
		})
	}

	tokenByte, err := jtoken.Parse(tokenString, func(jwtToken *jtoken.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jtoken.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return []byte("secret"), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(entity.Response{
			Code:    c.Response().StatusCode(),
			Message: fmt.Sprintf("invalidate token: %v", err),
			Data:    nil,
		})
	}

	_, ok := tokenByte.Claims.(jtoken.MapClaims)
	if !ok || !tokenByte.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(entity.Response{
			Code:    c.Response().StatusCode(),
			Message: "invalid token claim",
			Data:    nil,
		})

	}

	// var user models.User
	// initializers.DB.First(&user, "id = ?", fmt.Sprint(claims["sub"]))

	// if user.ID.String() != claims["sub"] {
	// 	return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "the user belonging to this token no logger exists"})
	// }

	// c.Locals("user", models.FilterUserRecord(&user))

	return c.Next()
}
