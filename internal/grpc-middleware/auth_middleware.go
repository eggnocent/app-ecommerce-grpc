package grpcmiddleware

import (
	"context"
	"fmt"
	jwtentity "github/eggnocent/app-grpc-eccomerce/internal/entity/jwt"
	"github/eggnocent/app-grpc-eccomerce/internal/utils"

	gocache "github.com/patrickmn/go-cache"

	"google.golang.org/grpc"
)

type AuthMiddleware struct {
	cacheService *gocache.Cache
}

var publicApi = map[string]bool{
	"/auth.AuthService/Login":                   true,
	"/auth.AuthService/Register":                true,
	"/product.ProductService/DetailProduct":     true,
	"/product.ProductService/ListProduct":       true,
	"/product.ProductService/HighlightProducts": true,
}

func (au *AuthMiddleware) Middleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	fmt.Println(info.FullMethod)
	if publicApi[info.FullMethod] {
		return handler(ctx, req)
	}
	// ambil token dari metadata
	tokenStr, err := jwtentity.ParseTokenFromContext(ctx)
	if err != nil {
		return nil, err
	}
	// cek token dari logout cache
	_, ok := au.cacheService.Get(tokenStr)
	if ok {
		return nil, utils.UnauthenticatedResponse()
	}
	// parse jwt hingga jadi entity
	claims, err := jwtentity.GetClaimsFromToken(tokenStr)
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, jwtentity.JwtEntityContextKeyValue, claims)

	// sematkan entity ke context
	ctx = claims.SetToContext(ctx)

	res, err := handler(ctx, req)

	return res, err
}

func NewAuthMiddleware(cacheService *gocache.Cache) *AuthMiddleware {
	return &AuthMiddleware{
		cacheService: cacheService,
	}
}
