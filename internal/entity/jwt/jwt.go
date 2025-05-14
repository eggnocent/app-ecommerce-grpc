package jwt

import (
	"context"
	"fmt"
	"github/eggnocent/app-grpc-eccomerce/internal/utils"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type JwtEntityContextKey string

var JwtEntityContextKeyValue JwtEntityContextKey = "JwtEntity"

type JwtClaims struct {
	jwt.RegisteredClaims
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Role     string `json:"role"`
}

func (jc *JwtClaims) SetToContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, JwtEntityContextKeyValue, jc)

}

func GetClaimsFromToken(token string) (*JwtClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(
		token,
		&JwtClaims{},
		func(t *jwt.Token) (interface{}, error) {
			if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Method.Alg())
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		},
	)

	if err != nil || !tokenClaims.Valid {
		return nil, utils.UnauthenticatedResponse()
	}

	claims, ok := tokenClaims.Claims.(*JwtClaims)

	if ok {
		return claims, nil
	}
	return nil, utils.UnauthenticatedResponse()

}
