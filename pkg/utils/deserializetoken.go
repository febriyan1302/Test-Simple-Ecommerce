package utils

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"

	jtoken "github.com/golang-jwt/jwt/v4"
)

type DeserializeTokenResponse struct {
	ID      int
	Email   string
	Expired int
}

func DeserializeToken(ctx *fiber.Ctx) DeserializeTokenResponse {
	var tokenString string
	authorization := ctx.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if ctx.Cookies("token") != "" {
		tokenString = ctx.Cookies("token")
	}

	tokenByte, _ := jtoken.Parse(tokenString, func(jwtToken *jtoken.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jtoken.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return []byte("secret"), nil
	})

	claims, _ := tokenByte.Claims.(jtoken.MapClaims)
	response := DeserializeTokenResponse{
		ID:      int(claims["ID"].(float64)),
		Email:   string(claims["email"].(string)),
		Expired: int(claims["exp"].(float64)),
	}

	return response
}
