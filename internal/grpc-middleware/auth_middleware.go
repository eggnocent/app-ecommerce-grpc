package grpcmiddleware

import (
	"context"
	jwtentity "github/eggnocent/app-grpc-eccomerce/internal/entity/jwt"
	"github/eggnocent/app-grpc-eccomerce/internal/utils"

	gocache "github.com/patrickmn/go-cache"

	"google.golang.org/grpc"
)

type AuthMiddleware struct {
	cacheService *gocache.Cache
}

func (au *AuthMiddleware) Middleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	if info.FullMethod == "/auth.AuthService/Login" || info.FullMethod == "/auth.AuthService/Register" {
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
